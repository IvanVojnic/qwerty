// Package handler declare handlers fo user
package handler

import (
	"net/http"

	"EFpractic2/models"
	"EFpractic2/pkg/utils"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type response struct {
	User *models.UserAuth `json:"user"`
}

func (h *Handler) signUp(c echo.Context) error {
	user := models.UserAuth{}
	errBind := c.Bind(&user)
	if errBind != nil {
		log.WithFields(log.Fields{
			"Error Bind json while creating user": errBind,
			"user":                                user,
		}).Info("Bind json")
		return echo.NewHTTPError(http.StatusInternalServerError, "data not correct")
	}
	user.UserID = uuid.New()
	rt, errRT := utils.GenerateToken(user.UserID, utils.TokenRTDuration)
	if errRT != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "smth went wrong")
	}
	at, errAT := utils.GenerateToken(user.UserID, utils.TokenATDuretion)
	if errAT != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "smth went wrong")
	}
	err := h.serviceProfile.CreateUserVerified(c.Request().Context(), &user, rt)
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

func (h *Handler) signIn(c echo.Context) error {
	user := models.UserAuth{}
	errBind := c.Bind(&user)
	if errBind != nil {
		log.WithFields(log.Fields{
			"Error Bind json while creating user": errBind,
			"user":                                user,
		}).Info("Bind json")
		return echo.NewHTTPError(http.StatusInternalServerError, "data not correct")
	}
	isVerified, err := h.serviceProfile.SignInUser(c.Request().Context(), &user)
	if err != nil {
		log.WithFields(log.Fields{
			"Error sign in user": err,
		}).Info("SIGN IN USER request")
		return echo.NewHTTPError(http.StatusBadRequest, "user sign in failed")
	}
	if isVerified {
		rt, errRT := utils.GenerateToken(user.UserID, utils.TokenRTDuration)
		if errRT != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "smth went wrong")
		}
		at, errAT := utils.GenerateToken(user.UserID, utils.TokenATDuretion)
		if errAT != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "smth went wrong")
		}
		return c.JSON(http.StatusOK, &Tokens{
			AccessToken:  at,
			RefreshToken: rt,
		})
	}
	return echo.NewHTTPError(http.StatusUnauthorized, "wrong data")
}

func (h *Handler) getUserAuth(c echo.Context) error {
	userID := c.Get("user_id").(uuid.UUID)
	user, err := h.serviceProfile.GetUserVerified(c.Request().Context(), userID)
	if err != nil {
		log.WithFields(log.Fields{
			"Error get user": err,
			"user":           user,
		}).Info("GET USER request")
		return echo.NewHTTPError(http.StatusBadRequest, "sign up please")
	}
	return c.JSON(http.StatusOK, response{&user})
}

func (h *Handler) refreshToken(c echo.Context) error {
	var tokens Tokens
	errBind := c.Bind(&tokens)
	if errBind != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "cannot bind")
	}
	checkATexpired := utils.IsTokenExpired(tokens.AccessToken)
	if checkATexpired {
		checkRT, errRT := utils.ParseToken(tokens.RefreshToken)
		if checkRT {
			if errRT != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "bad refresh token")
			}
			id, errGetID := utils.ExtractIDFromToken(tokens.RefreshToken)
			if errGetID != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "smth went wrong")
			}
			newAt, errAT := utils.GenerateToken(id, utils.TokenATDuretion)
			if errAT != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "cannot bind")
			}
			return c.JSON(http.StatusOK, Tokens{AccessToken: newAt, RefreshToken: tokens.RefreshToken})
		}
		return echo.NewHTTPError(http.StatusUnauthorized, "login please")
	}
	return echo.NewHTTPError(http.StatusUnauthorized, "login please")
}
