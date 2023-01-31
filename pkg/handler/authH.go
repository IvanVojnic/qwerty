package handler

import (
	"EFpractic2/models"
	"EFpractic2/pkg/utils"
	"context"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Tokens struct {
	AccessToken  string `json:"access"`
	RefreshToken string `json:"refresh"`
}

type Response struct {
	Tokens *Tokens          `json:"tokens"`
	User   *models.UserAuth `json:"user"`
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

func (h *Handler) GetIdByAT(ctx context.Context, accessToken string, refreshToken string) (int, string, error) {
	id, err := utils.ParseToken(accessToken)
	if err != nil {
		log.WithFields(log.Fields{
			"Error parse token": err,
			"id":                id,
		}).Info("Parse access token")
		accessToken, err = h.RefreshAccessToken(refreshToken)
		if err != nil {
			return 0, "", err
		}
		h.GetIdByAT(ctx, accessToken, refreshToken)
	}
	return id, accessToken, nil
}

func (h *Handler) RefreshAccessToken(refreshToken string) (string, error) {
	id, err := utils.ParseToken(refreshToken)
	if err != nil {
		log.WithFields(log.Fields{
			"Error parse token": err,
			"id":                id,
		}).Info("Parse refresh token")
		return "", err
	}
	newAccessToken, err := utils.GenerateToken(id, false)
	if err != nil {
		log.WithFields(log.Fields{
			"Error parse token": err,
			"id":                id,
		}).Info("Parse refresh token")
		return "", err
	}
	return newAccessToken, nil
}

func (h *Handler) GetUserAuth(c echo.Context) error {
	var tokens Tokens
	err := c.Bind(&tokens)
	id, accessToken, err := h.GetIdByAT(context.Background(), tokens.AccessToken, tokens.RefreshToken)
	if err != nil {
		log.WithFields(log.Fields{
			"Error Bind json while creating user": err,
			"tokens":                              tokens,
		}).Info("Bind json")
		return echo.NewHTTPError(http.StatusInternalServerError, "data not correct")
	}
	//user, _, err := h.services.Authorization.GetUserVerified(c.Request().Context(), tokens.AccessToken, tokens.RefreshToken)
	user, err := h.services.Authorization.GetUserVerified(c.Request().Context(), id)
	if err != nil {
		log.WithFields(log.Fields{
			"Error get user": err,
			"user":           user,
		}).Info("GET USER request")
		return echo.NewHTTPError(http.StatusBadRequest, "sign up please")
	}
	return c.JSON(http.StatusOK, Response{
		&Tokens{
			AccessToken: accessToken, RefreshToken: tokens.RefreshToken,
		},
		&user,
	})
}
