package handler

import (
	"EFpractic2/models"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Tokens struct {
	AccessToken  string `json:"access"`
	RefreshToken string `json:"refresh"`
}

func (h *Handler) SignUp(c echo.Context) error {
	user := models.UserAuth{}
	err := c.Bind(&user)
	if err != nil {
		log.WithFields(log.Fields{
			"Error Bind json while creating user": err,
			"user":                                user,
		}).Info("Bind json")
		return echo.NewHTTPError(http.StatusInternalServerError, "data not correct")
	}
	rt, at, err := h.services.Authorization.CreateUserVerified(c.Request().Context(), user)
	if err != nil {
		log.WithFields(log.Fields{
			"Error create user": err,
			"access token":      at,
			"refresh token":     rt,
		}).Info("CREATE USER request")
		return echo.NewHTTPError(http.StatusBadRequest, "user creating failed")
	}
	return c.JSON(http.StatusOK, &Tokens{
		AccessToken:  at,
		RefreshToken: rt,
	})
}

func (h *Handler) GetUserAuth(c echo.Context) error {
	var tokens Tokens
	err := c.Bind(&tokens)
	if err != nil {
		log.WithFields(log.Fields{
			"Error Bind json while creating user": err,
			"tokens":                              tokens,
		}).Info("Bind json")
		return echo.NewHTTPError(http.StatusInternalServerError, "data not correct")
	}
	user, _, err := h.services.Authorization.GetUserVerified(c.Request().Context(), tokens.AccessToken, tokens.RefreshToken)
	if err != nil {
		log.WithFields(log.Fields{
			"Error get user": err,
			"user":           user,
		}).Info("GET USER request")
		return echo.NewHTTPError(http.StatusBadRequest, "sign up please")
	}
	return c.JSON(http.StatusOK, user)
}
