package profile

type ProfileService struct {
	profileRepository ProfileRepository
}

func NewProfileService(profileRepository ProfileRepository) ProfileService {
	return ProfileService{profileRepository}
}

func (p *ProfileService) Create(input Profile) (Profile, error) {
	// TODO: hash password
	return p.profileRepository.Create(input)
}

func (p *ProfileService) Get(id ProfileID) (Profile, error) {
	return p.profileRepository.Get(id)
}
