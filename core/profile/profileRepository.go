package profile

type ProfileRepository interface {
	Create(input Profile) (Profile, error)
	Get(id ProfileID) (Profile, error)
	Update(id ProfileID, input Profile) (Profile, error)
	Delete(id ProfileID) error
}
