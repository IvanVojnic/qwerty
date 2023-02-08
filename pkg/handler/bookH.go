// Package handler declare handlers for book
package handler

import (
	"EFpractic2/models"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) HomeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"name": "upload",
	})
}

// CreateBook used to create book
func (h *Handler) CreateBook(c echo.Context) error { // nolint:dupl, gocritic
	book := models.Book{}
	err := c.Bind(&book)
	if err != nil {
		log.WithFields(log.Fields{
			"Error Bind json while creating book": err,
			"book":                                book,
		}).Info("Bind json")
		return echo.NewHTTPError(http.StatusInternalServerError, "data not correct")
	}
	err = h.serviceBook.CreateBook(c.Request().Context(), book)
	if err != nil {
		log.WithFields(log.Fields{
			"Error create book": err,
			"book":              book,
		}).Info("CREATE BOOK request")
		return echo.NewHTTPError(http.StatusBadRequest, "book creating failed")
	}
	return c.String(http.StatusOK, "book created")
}

// UpdateBook used to update book
func (h *Handler) UpdateBook(c echo.Context) error { // nolint:dupl, gocritic
	book := models.Book{}
	err := c.Bind(&book)
	if err != nil {
		log.WithFields(log.Fields{
			"Error Bind json while updating book": err,
			"book":                                book,
		}).Info("Bind json")
		return echo.NewHTTPError(http.StatusInternalServerError, "data not correct")
	}
	err = h.serviceBook.UpdateBook(c.Request().Context(), book)
	if err != nil {
		log.WithFields(log.Fields{
			"Error update book": err,
			"book":              book,
		}).Info("UPDATE BOOK request")
		return echo.NewHTTPError(http.StatusBadRequest, "book updating failed")
	}
	return c.String(http.StatusOK, "book updated")
}

// DeleteBook used to delete book
func (h *Handler) DeleteBook(c echo.Context) error {
	bookName := c.QueryParam("name")
	err := h.serviceBook.DeleteBook(c.Request().Context(), bookName)
	if err != nil {
		log.WithFields(log.Fields{
			"Error get book": err,
			"book ID":        bookName,
		}).Info("DELETE BOOK request")
		return echo.NewHTTPError(http.StatusBadRequest, "book deleting failed")
	}
	return c.String(http.StatusOK, "bool deleted")
}

// GetAllBooks used to get all books
func (h *Handler) GetAllBooks(c echo.Context) error {
	books, err := h.serviceBook.GetAllBooks(c.Request().Context())
	if err != nil {
		log.WithFields(log.Fields{
			"Error get all books": err,
			"books":               books,
		}).Info("GET ALL BOOK request")
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, books)
}

// GetBookByName used to get book by name
func (h *Handler) GetBookByName(c echo.Context) error {
	bookName := c.QueryParam("name")
	book, err := h.serviceBook.GetBookByName(c.Request().Context(), bookName)
	if err != nil {
		log.WithFields(log.Fields{
			"Error get book": err,
			"book":           book,
		}).Info("GET BOOK request")
		return echo.NewHTTPError(http.StatusBadRequest, "cannot get book")
	}
	return c.JSON(http.StatusOK, book)
}
