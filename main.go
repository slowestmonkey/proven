package main

import (
	"database/sql"
	"fmt"
	"log"
	profileRepository "proven/adapters/database/profile"
	profileHttp "proven/adapters/http/profile"
	profile "proven/core/profile"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	connection := viper.GetString(`database.connection`)
	fmt.Println(connection)
	db, err := sql.Open("postgres", connection)

	if err != nil {
		log.Fatal(err)
	}

	profileRepo := profileRepository.NewProfileRepository(db)
	profileService := profile.NewProfileService(profileRepo)
	profileHttp.NewProfileHandler(e, profileService)

	e.Logger.Fatal(e.Start(":3000"))
}
