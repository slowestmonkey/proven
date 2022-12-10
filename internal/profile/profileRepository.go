package profile

type ProfileRepository interface {
	Store(input Profile) (Profile, error)
	Get(id string) (Profile, error)
	Update(id string, input Profile) error
	Archive(id string) error
}
