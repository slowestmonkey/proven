package app

import (
	"database/sql"
	"proven/internal/adapters/database"
	"proven/internal/adapters/http"
	"proven/internal/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func Run() {
	e := echo.New()
	e.Use(middleware.Logger())

	connection := viper.GetString(`database.connection`)
	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err)
	}

	profileRepo := database.NewProfileRepository(db)
	profileUseCase := usecase.NewProfileUseCase(profileRepo)
	http.NewProfileHandler(e, profileUseCase)

	e.Logger.Fatal(e.Start(viper.GetString(`server.address`)))
}
