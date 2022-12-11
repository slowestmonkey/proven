package entity

import (
	"time"
)

type Email string // TODO replace with struct{ email string, confirmed boolean }

type Country string

type Citizenship string

type PhoneNumber string

type HashedPassword string

type Profile struct {
	ID               string         `json:"id"`
	FirstName        string         `json:"firstName"`
	LastName         string         `json:"lastName"`
	Email            Email          `json:"email"`
	PhoneNumber      PhoneNumber    `json:"phoneNumber"`
	Citizenship      Citizenship    `json:"citizenship"`
	BirthDate        time.Time      `json:"birthDate"`
	BirthCountry     Country        `json:"birthCountry"`
	ResidenceCountry Country        `json:"residenceCountry"`
	Password         HashedPassword `json:"-"`
	Account          Account
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"-"`
	ArchivedAt       time.Time `json:"-"`
}

type ProfileRepository interface {
	Store(input Profile) (Profile, error)
	Get(id string) (Profile, error)
	Update(id string, input Profile) error
	Archive(id string) error
}
