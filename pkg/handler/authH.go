package handler

import (
	"EFpractic2/models"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) createVerifiedUser(c echo.Context) error {
	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		log.WithFields(log.Fields{
			"Error Bind json while creating user": err,
			"user":                                user,
		}).Info("Bind json")
		return echo.NewHTTPError(http.StatusInternalServerError, "data not correct")
	}
	err = h.services.UserAct.CreateUser(c.Request().Context(), user)
	if err != nil {
		log.WithFields(log.Fields{
			"Error create user": err,
			"user":              user,
		}).Info("CREATE USER request")
		return echo.NewHTTPError(http.StatusBadRequest, "user creating failed")
	}
	return c.String(http.StatusOK, "user created")
}
