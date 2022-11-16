package main

import (
	profileRepository "proven/adapters/database/profile"
	profileHttp "proven/adapters/http/profile"
	profile "proven/core/profile"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	profileRepo := profileRepository.NewProfileRepository("")
	profileService := profile.NewProfileService(profileRepo)
	profileHttp.NewProfileHandler(e, profileService)

	e.Logger.Fatal(e.Start(":3000"))
}
