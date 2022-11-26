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
		INSERT INTO profile (first_name, last_name, email, phone_number, citizenship, birth_date, birth_country, residence_country, password)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
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
	return profile.Profile{}, nil
}

func (p *ProfileRepository) Update(id string, input profile.Profile) (profile.Profile, error) {
	return profile.Profile{}, nil
}

func (p *ProfileRepository) Delete(id string) error {
	return nil
}
