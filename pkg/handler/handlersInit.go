// Package handler init all handlers and middleware
package handler

import (
	"context"
	"fmt"
	"net/http"

	"EFpractic2/models"
	"EFpractic2/pkg/errorwrapper"
	"EFpractic2/pkg/utils"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// BookAct service consists of methods fo book
type BookAct interface {
	CreateBook(context.Context, models.Book) error
	UpdateBook(context.Context, models.Book) error
	GetBook(context.Context, int) (models.Book, error)
	DeleteBook(context.Context, int) error
	GetAllBooks(context.Context) ([]models.Book, error)
}

// Authorization service consists of methods fo user
type Authorization interface {
	CreateUserVerified(context.Context, *models.UserAuth, string) error
	GetUserVerified(context.Context, uuid.UUID) (models.UserAuth, error)
	SignInUser(context.Context, *models.UserAuth) (bool, error)
}

// Handler is wrapper for service
type Handler struct {
	serviceProfile Authorization
	serviceBook    BookAct
}

// NewHandler used to init Handler
func NewHandler(serviceAuth Authorization, serviceBook BookAct) *Handler {
	return &Handler{
		serviceProfile: serviceAuth,
		serviceBook:    serviceBook,
	}
}

// Tokens used to define at and rt
type Tokens struct {
	AccessToken  string `json:"access"`
	RefreshToken string `json:"refresh"`
}

// InitRoutes used to init routes
func (h *Handler) InitRoutes(router *echo.Echo) *echo.Echo {
	router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})
	rAct := router.Group("/book")
	router.POST("/refreshToken", h.refreshToken)
	router.POST("/createUser", h.signUp)
	router.POST("/signIn", h.signIn)
	rAct.Use(middleware.Logger())
	rAct.POST("/create", h.createBook)
	rAct.GET("/get", h.getBook)
	rAct.POST("/update", h.updateBook)
	rAct.GET("/delete", h.deleteBook)
	rAct.GET("/getAllBooks", h.getAllBooks)
	rVerified := router.Group("/verified")
	rVerified.Use(jwtAuthMiddleware())
	rVerified.POST("/getUserAuth", h.getUserAuth)
	router.Logger.Fatal(router.Start(":40000"))
	return router
}

func jwtAuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			var tokens Tokens
			err = c.Bind(&tokens)
			if err != nil {
				_ = fmt.Errorf("error after binding tokens %w", err)
				return echo.NewHTTPError(http.StatusUnauthorized, "login please")
			}
			authorized, errIsAuth := utils.IsAuthorized(tokens.AccessToken)
			if authorized {
				userID, errGetID := utils.ExtractIDFromToken(tokens.AccessToken)
				if errGetID != nil {
					return echo.NewHTTPError(http.StatusUnauthorized, errorwrapper.ErrorResponse{Message: errGetID.Error()})
				}
				c.Set("user_id", userID)
				return next(c)
			}
			return echo.NewHTTPError(http.StatusUnauthorized, errorwrapper.ErrorResponse{Message: errIsAuth.Error()})
		}
	}
}

/*func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}*/
