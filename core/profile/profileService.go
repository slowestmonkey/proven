package profile

import (
	"proven/core/account"

	"github.com/google/uuid"
)

type ProfileService struct {
	profileRepository ProfileRepository
}

func NewProfileService(profileRepository ProfileRepository) ProfileService {
	return ProfileService{profileRepository}
}

func (p *ProfileService) Create(input ProfileCreateInput) (Profile, error) {
	ID := uuid.New()
	hashedPassword := input.Password // TODO: hash password

	return p.profileRepository.Create(Profile{
		ID:               ProfileID(ID),
		FirstName:        input.FirstName,
		LastName:         input.LastName,
		Email:            Email(input.Email),
		PhoneNumber:      PhoneNumber(input.PhoneNumber),
		Citizenship:      Citizenship(input.Citizenship),
		BirthDate:        input.BirthDate,
		BirthCountry:     Country(input.BirthCountry),
		ResidenceCountry: Country(input.ResidenceCountry),
		HashedPassword:   HashedPassword(hashedPassword),
		Account:          account.Account{},
	})
}

func (p *ProfileService) Get(id ProfileID) (Profile, error) {
	return p.profileRepository.Get(id)
}
