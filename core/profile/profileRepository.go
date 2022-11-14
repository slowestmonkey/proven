package profile

type ProfileRepository interface {
	Create(input Profile) (Profile, error)
	Get(id ProfileID) (Profile, error)
	Update(id ProfileID, input ProfileUpdateInput) (Profile, error)
	Delete(id ProfileID) error
}
