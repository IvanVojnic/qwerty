// Package handler declare handlers for book
package handler

import (
	"net/http"
	"strconv"

	"EFpractic2/models"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) createBook(c echo.Context) error { // nolint:dupl, gocritic
	book := models.Book{}
	err := c.Bind(&book)
	if err != nil {
		log.WithFields(log.Fields{
			"Error Bind json while creating book": err,
			"book":                                book,
		}).Info("Bind json")
		return echo.NewHTTPError(http.StatusInternalServerError, "data not correct")
	}
	err = h.services.BookAct.CreateBook(c.Request().Context(), book)
	if err != nil {
		log.WithFields(log.Fields{
			"Error create book": err,
			"book":              book,
		}).Info("CREATE BOOK request")
		return echo.NewHTTPError(http.StatusBadRequest, "book creating failed")
	}
	return c.String(http.StatusOK, "book created")
}

func (h *Handler) getBook(c echo.Context) error {
	bookID := c.QueryParam("id")
	var bookIDNum int
	bookIDNum, _ = strconv.Atoi(bookID)
	book, err := h.services.BookAct.GetBook(c.Request().Context(), bookIDNum)
	if err != nil {
		log.WithFields(log.Fields{
			"Error get book": err,
			"book":           book,
		}).Info("GET BOOK request")
		return echo.NewHTTPError(http.StatusBadRequest, "book getting failed")
	}
	return c.JSON(http.StatusOK, book)
}

func (h *Handler) updateBook(c echo.Context) error { // nolint:dupl, gocritic
	book := models.Book{}
	err := c.Bind(&book)
	if err != nil {
		log.WithFields(log.Fields{
			"Error Bind json while updating book": err,
			"book":                                book,
		}).Info("Bind json")
		return echo.NewHTTPError(http.StatusInternalServerError, "data not correct")
	}
	err = h.services.BookAct.UpdateBook(c.Request().Context(), book)
	if err != nil {
		log.WithFields(log.Fields{
			"Error update book": err,
			"book":              book,
		}).Info("UPDATE BOOK request")
		return echo.NewHTTPError(http.StatusBadRequest, "book updating failed")
	}
	return c.String(http.StatusOK, "book updated")
}

func (h *Handler) deleteBook(c echo.Context) error {
	bookID := c.QueryParam("id")
	var bookIDNum int
	bookIDNum, _ = strconv.Atoi(bookID)
	err := h.services.BookAct.DeleteBook(c.Request().Context(), bookIDNum)
	if err != nil {
		log.WithFields(log.Fields{
			"Error get book": err,
			"book ID":        bookID,
		}).Info("DELETE BOOK request")
		return echo.NewHTTPError(http.StatusBadRequest, "book deleting failed")
	}
	return c.String(http.StatusOK, "bool deleted")
}

func (h *Handler) getAllBooks(c echo.Context) error {
	books, err := h.services.BookAct.GetAllBooks(c.Request().Context())
	if err != nil {
		log.WithFields(log.Fields{
			"Error get all books": err,
			"books":               books,
		}).Info("GET ALL BOOK request")
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, books)
}
