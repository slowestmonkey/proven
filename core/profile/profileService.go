package profile

type ProfileService struct {
	profileRepository ProfileRepository
}

func NewProfileService(profileRepository ProfileRepository) *ProfileService {
	return &ProfileService{profileRepository}
}

func (p *ProfileService) Store(input Profile) (Profile, error) {
	// TODO: hash password
	return p.profileRepository.Store(input)
}

func (p *ProfileService) Get(id string) (Profile, error) {
	return p.profileRepository.Get(id)
}
