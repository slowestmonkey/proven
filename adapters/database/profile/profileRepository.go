package profilerepo

import (
	"database/sql"
	profile "proven/core/profile"
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
		return profile.Profile{}, err
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

	if err != nil {
		return profile.Profile{}, err
	}

	return result, nil
}

func (p *ProfileRepository) Update(id string, input profile.Profile) (profile.Profile, error) {
	return profile.Profile{}, nil
}

func (p *ProfileRepository) Delete(id string) error {
	return nil
}
