package profilehttp

import (
	"net/http"
	profile "proven/core/profile"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ProfileHandler struct {
	service *profile.ProfileService
}

func NewProfileHandler(e *echo.Echo, service *profile.ProfileService) {
	handler := ProfileHandler{service}

	e.GET("/profiles/:id", handler.Fetch)
	e.POST("/profiles", handler.Store)
	e.PATCH("/profiles/:id", handler.Update)
	e.DELETE("/profiles/:id", handler.Archive)
}

func (p *ProfileHandler) Store(ctx echo.Context) error {
	var input profile.Profile

	err := ctx.Bind(&input)

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isProfileInputValid(input); !ok {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	// context := ctx.Request().Context()
	profile, err := p.service.Store(input)

	if err != nil {
		// TODO: should check the domain error
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusCreated, profile)
}

func isProfileInputValid(m profile.Profile) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)

	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *ProfileHandler) Fetch(ctx echo.Context) error {
	profile, err := p.service.Get(ctx.Param("id"))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, profile)
}

func (p *ProfileHandler) Update(ctx echo.Context) error {
	var input profile.Profile

	err := ctx.Bind(&input)

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isProfileInputValid(input); !ok {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err = p.service.Update(ctx.Param("id"), input)

	if err != nil {
		// TODO: should check the domain error
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}

func (p *ProfileHandler) Archive(ctx echo.Context) error {
	err := p.service.Archive(ctx.Param("id"))

	if err != nil {
		// TODO: might be also 404
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}
