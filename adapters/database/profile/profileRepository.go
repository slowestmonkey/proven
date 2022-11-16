package profilerepo

import (
	"proven/core/account"
	profile "proven/core/profile"

	"github.com/google/uuid"
)

type ProfileRepository struct {
	connection string
}

func NewProfileRepository(connection string) profile.ProfileRepository {
	return &ProfileRepository{connection}
}

func (p *ProfileRepository) Create(input profile.Profile) (profile.Profile, error) {
	return profile.Profile{
		ID:               profile.ProfileID(uuid.New()),
		FirstName:        input.FirstName,
		LastName:         input.LastName,
		Email:            profile.Email(input.Email),
		PhoneNumber:      profile.PhoneNumber(input.PhoneNumber),
		Citizenship:      profile.Citizenship(input.Citizenship),
		BirthDate:        input.BirthDate,
		BirthCountry:     profile.Country(input.BirthCountry),
		ResidenceCountry: profile.Country(input.ResidenceCountry),
		Password:         input.Password,
		Account:          account.Account{},
	}, nil
}

func (p *ProfileRepository) Get(id profile.ProfileID) (profile.Profile, error) {
	return profile.Profile{}, nil
}

func (p *ProfileRepository) Update(id profile.ProfileID, input profile.Profile) (profile.Profile, error) {
	return profile.Profile{}, nil
}

func (p *ProfileRepository) Delete(id profile.ProfileID) error {
	return nil
}
