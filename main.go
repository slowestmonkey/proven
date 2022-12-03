package main

import (
	"database/sql"
	database "proven/adapters/database"
	http "proven/adapters/http"
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
	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err)
	}

	profileRepo := database.NewProfileRepository(db)
	profileUseCase := profile.New(profileRepo)
	http.NewProfileHandler(e, profileUseCase)

	e.Logger.Fatal(e.Start(":3000"))
}
