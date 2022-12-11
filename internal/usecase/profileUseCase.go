package usecase

import (
	"proven/internal/entity"

	"golang.org/x/crypto/bcrypt"
)

type ProfileUseCase struct {
	profileRepository entity.ProfileRepository
}

func NewProfileUseCase(profileRepository entity.ProfileRepository) *ProfileUseCase {
	return &ProfileUseCase{profileRepository}
}

func (uc *ProfileUseCase) Store(input entity.Profile) (entity.Profile, error) {
	hashedPassword, err := hashPassword(string(input.Password))

	if err != nil {
		return entity.Profile{}, err
	}

	input.Password = hashedPassword

	return uc.profileRepository.Store(input)
}

func hashPassword(password string) (entity.HashedPassword, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	return entity.HashedPassword(string(bytes)), err
}

func (uc *ProfileUseCase) Get(id string) (entity.Profile, error) {
	return uc.profileRepository.Get(id)
}

func (uc *ProfileUseCase) Update(id string, input entity.Profile) error {
	return uc.profileRepository.Update(id, input)
}

func (uc *ProfileUseCase) Archive(id string) error {
	return uc.profileRepository.Archive(id)
}
