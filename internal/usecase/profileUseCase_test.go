package usecase_test

import (
	"database/sql"
	"os"
	"reflect"
	"testing"
	"time"

	"proven/internal/adapters/database"
	"proven/internal/entity"
	"proven/internal/usecase"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
)

// TODO: fix relative paths
// TODO: test password hashing

var (
	db *sql.DB
	uc usecase.ProfileUseCase
	m  *migrate.Migrate
)

func setup() {
	viper.SetConfigFile("../../config/config.test.json")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	connection := viper.GetString(`database.connection`)
	db, err = sql.Open("postgres", connection)

	if err != nil {
		panic(err)
	}

	m, err = migrate.New(
		"file://../../migrations",
		connection,
	)

	if err != nil {
		panic(err)
	}

	m.Up()

	profileRepo := database.NewProfileRepository(db)
	uc = *usecase.NewProfileUseCase(profileRepo)
}

func shutdown() {
	m.Down()
	db.Close()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func TestCreate(t *testing.T) {
	input := entity.Profile{
		FirstName:        "Jamie",
		LastName:         "Foxx",
		Email:            "jamie@foxx.com",
		PhoneNumber:      "5124125",
		Citizenship:      "US",
		BirthDate:        time.Date(1988, time.Month(2), 21, 1, 10, 30, 0, time.UTC),
		BirthCountry:     "US",
		ResidenceCountry: "US",
	}

	stored, err := uc.Store(input)
	if err != nil {
		t.Errorf(err.Error())
	}

	expected := entity.Profile{
		FirstName:        stored.FirstName,
		LastName:         stored.LastName,
		Email:            stored.Email,
		PhoneNumber:      stored.PhoneNumber,
		Citizenship:      stored.Citizenship,
		BirthDate:        stored.BirthDate,
		BirthCountry:     stored.BirthCountry,
		ResidenceCountry: stored.ResidenceCountry,
	}
	if !reflect.DeepEqual(input, expected) {
		t.Error("Cannot store profile")
	}
}

func TestFetch(t *testing.T) {
	input := entity.Profile{
		Email: "jamie@foxx.com",
	}

	stored, err := uc.Store(input)
	if err != nil {
		t.Errorf(err.Error())
	}

	fetched, err := uc.Get(stored.ID)
	if err != nil {
		t.Errorf(err.Error())
	}

	if input.Email != fetched.Email {
		t.Error("Cannot fetch profile")
	}
}

func TestUpdate(t *testing.T) {
	stored, _ := uc.Store(entity.Profile{
		FirstName:        "Jamie",
		LastName:         "Foxx",
		Email:            "jamie@foxx.com",
		PhoneNumber:      "5124125",
		Citizenship:      "US",
		BirthDate:        time.Date(1988, time.Month(2), 21, 1, 10, 30, 0, time.UTC),
		BirthCountry:     "US",
		ResidenceCountry: "US",
	})
	input := entity.Profile{
		FirstName:        "Hue",
		LastName:         "Jackman",
		PhoneNumber:      "1234567",
		Citizenship:      "AU",
		ResidenceCountry: "AU",
	}

	err := uc.Update(stored.ID, input)
	if err != nil {
		t.Errorf(err.Error())
	}

	fetched, _ := uc.Get(stored.ID)

	expected := entity.Profile{
		FirstName:        fetched.FirstName,
		LastName:         fetched.LastName,
		PhoneNumber:      fetched.PhoneNumber,
		Citizenship:      fetched.Citizenship,
		ResidenceCountry: fetched.ResidenceCountry,
	}
	if !reflect.DeepEqual(input, expected) {
		t.Error("Cannot update profile")
	}
}

func TestArchive(t *testing.T) {
	input := entity.Profile{
		Email: "jamie@foxx.com",
	}

	stored, err := uc.Store(input)
	if err != nil {
		t.Errorf(err.Error())
	}

	err = uc.Archive(stored.ID)
	if err != nil {
		t.Errorf(err.Error())
	}

	archived, _ := uc.Get(stored.ID)
	if archived.ArchivedAt.IsZero() {
		t.Errorf("Cannot archive profile")
	}
}
