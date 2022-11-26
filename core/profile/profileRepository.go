package profile

type ProfileRepository interface {
	Store(input Profile) (Profile, error)
	Get(id string) (Profile, error)
	Update(id string, input Profile) (Profile, error)
	Delete(id string) error
}
