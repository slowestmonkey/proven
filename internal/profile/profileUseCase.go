package profile

import "golang.org/x/crypto/bcrypt"

type ProfileUseCase struct {
	profileRepository ProfileRepository
}

func New(profileRepository ProfileRepository) *ProfileUseCase {
	return &ProfileUseCase{profileRepository}
}

func (uc *ProfileUseCase) Store(input Profile) (Profile, error) {
	hashedPassword, err := hashPassword(string(input.Password))

	if err != nil {
		return Profile{}, err
	}

	input.Password = hashedPassword

	return uc.profileRepository.Store(input)
}

func hashPassword(password string) (HashedPassword, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	return HashedPassword(string(bytes)), err
}

func (uc *ProfileUseCase) Get(id string) (Profile, error) {
	return uc.profileRepository.Get(id)
}

func (uc *ProfileUseCase) Update(id string, input Profile) error {
	return uc.profileRepository.Update(id, input)
}

func (uc *ProfileUseCase) Archive(id string) error {
	return uc.profileRepository.Archive(id)
}
