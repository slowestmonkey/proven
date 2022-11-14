package profile

import (
	"proven/core/account"
	"time"

	"github.com/google/uuid"
)

type ProfileID uuid.UUID

type Email string // TODO replace with struct{ email string, confirmed boolean }

type Country string

type Citizenship string

type PhoneNumber string

type HashedPassword string

type ProfileCreateInput struct {
	FirstName        string    `json:"firstName"`
	LastName         string    `json:"lastName"`
	Email            string    `json:"email"`
	PhoneNumber      string    `json:"phoneNumber"`
	Citizenship      string    `json:"citizenship"`
	BirthDate        time.Time `json:"birthDate"`
	BirthCountry     string    `json:"birthCountry"`
	ResidenceCountry string    `json:"residenceCountry"`
	Password         string    `json:"password"`
}

type ProfileUpdateInput struct {
	PhoneNumber      string `json:"phoneNumber"`
	ResidenceCountry string `json:"residenceCountry"`
}

type Profile struct {
	ID               ProfileID      `json:"id"`
	FirstName        string         `json:"firstName"`
	LastName         string         `json:"lastName"`
	Email            Email          `json:"email"`
	PhoneNumber      PhoneNumber    `json:"phoneNumber"`
	Citizenship      Citizenship    `json:"citizenship"`
	BirthDate        time.Time      `json:"birthDate"`
	BirthCountry     Country        `json:"birthCountry"`
	ResidenceCountry Country        `json:"residenceCountry"`
	HashedPassword   HashedPassword `json:"hashedPassword"`
	Account          account.Account
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
	ArchivedAt       time.Time `json:"ArchivedAt"`
}
