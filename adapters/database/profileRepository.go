package database

import (
	"database/sql"
	domainerr "proven/core/error"
	profile "proven/core/profile"
	"time"
)

type ProfileRepository struct {
	db *sql.DB
}

func NewProfileRepository(db *sql.DB) profile.ProfileRepository {
	return &ProfileRepository{db}
}

func (p *ProfileRepository) Store(input profile.Profile) (profile.Profile, error) {
	query := `
		INSERT INTO profile (first_name, last_name, email, phone_number, citizenship, birth_date, birth_country, residence_country, password, updated_at, archived_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id, created_at
	`

	err := p.db.QueryRow(query,
		input.FirstName,
		input.LastName,
		input.Email,
		input.PhoneNumber,
		input.Citizenship,
		input.BirthDate,
		input.BirthCountry,
		input.ResidenceCountry,
		input.Password,
		input.UpdatedAt,
		input.ArchivedAt,
	).Scan(
		&input.ID,
		&input.CreatedAt,
	)

	if err != nil {
		return profile.Profile{}, domainerr.ErrInternalServerError
	}
	return input, nil
}

func (p *ProfileRepository) Get(id string) (profile.Profile, error) {
	query := `SELECT * FROM profile WHERE id = $1`

	var result profile.Profile

	err := p.db.QueryRow(query, id).Scan(
		&result.ID,
		&result.FirstName,
		&result.LastName,
		&result.Email,
		&result.PhoneNumber,
		&result.Citizenship,
		&result.BirthDate,
		&result.BirthCountry,
		&result.ResidenceCountry,
		&result.Password,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.ArchivedAt,
	)

	switch err {
	case nil:
		return result, nil
	case sql.ErrNoRows:
		return profile.Profile{}, domainerr.ErrNotFound
	default:
		return profile.Profile{}, domainerr.ErrInternalServerError
	}
}

func (p *ProfileRepository) Update(id string, input profile.Profile) error {
	query := `
		UPDATE profile SET first_name=$1, last_name=$2, phone_number=$3, citizenship=$4, residence_country=$5, updated_at=$6
		WHERE id = $7
	`

	_, err := p.db.Exec(
		query,
		&input.FirstName,
		&input.LastName,
		&input.PhoneNumber,
		&input.Citizenship,
		&input.ResidenceCountry,
		time.Now(),
		id,
	)

	if err != nil {
		return domainerr.ErrInternalServerError
	}
	return nil
}

func (p *ProfileRepository) Archive(id string) error {
	query := `UPDATE profile SET archived_at=$1 WHERE id = $2`

	_, err := p.db.Exec(query, time.Now(), id)

	if err != nil {
		return domainerr.ErrInternalServerError
	}
	return nil
}
