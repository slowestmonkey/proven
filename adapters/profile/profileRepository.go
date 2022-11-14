package profilerepo

import (
	"proven/core/profile"
)

type ProfileRepository struct {
	connection string
}

func NewProfileRepository(connection string) profile.ProfileRepository {
	return &ProfileRepository{connection}
}

func (p *ProfileRepository) Create(input profile.Profile) (profile.Profile, error) {
	return input, nil
}

func (p *ProfileRepository) Get(id profile.ProfileID) (profile.Profile, error) {
	return profile.Profile{}, nil
}

func (p *ProfileRepository) Update(id profile.ProfileID, input profile.ProfileUpdateInput) (profile.Profile, error) {
	return profile.Profile{}, nil
}

func (p *ProfileRepository) Delete(id profile.ProfileID) error {
	return nil
}
