package profile

import "golang.org/x/crypto/bcrypt"

type ProfileService struct {
	profileRepository ProfileRepository
}

func NewProfileService(profileRepository ProfileRepository) *ProfileService {
	return &ProfileService{profileRepository}
}

func (p *ProfileService) Store(input Profile) (Profile, error) {
	hashedPassword, err := hashPassword(string(input.Password))

	if err != nil {
		return Profile{}, err
	}

	input.Password = hashedPassword

	return p.profileRepository.Store(input)
}

func hashPassword(password string) (HashedPassword, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	return HashedPassword(string(bytes)), err
}

func (p *ProfileService) Get(id string) (Profile, error) {
	return p.profileRepository.Get(id)
}

func (p *ProfileService) Update(id string, input Profile) error {
	return p.profileRepository.Update(id, input)
}

func (p *ProfileService) Archive(id string) error {
	return p.profileRepository.Archive(id)
}
