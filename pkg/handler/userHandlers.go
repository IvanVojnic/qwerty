package handler

import (
	"EFpractic2/models"
	"fmt"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) createBook(c echo.Context) error {
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
	bookId := c.QueryParam("id")
	var bookIdNum int
	bookIdNum, _ = strconv.Atoi(bookId)
	book, err := h.services.BookAct.GetBook(c.Request().Context(), bookIdNum)
	if err != nil {
		log.WithFields(log.Fields{
			"Error get book": err,
			"book":           book,
		}).Info("GET BOOK request")
		return echo.NewHTTPError(http.StatusBadRequest, "book getting failed")
	}
	fmt.Sprintf("book: %s", book)
	return c.JSON(http.StatusOK, book)
}

func (h *Handler) updateBook(c echo.Context) error {
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
	bookId := c.QueryParam("id")
	var bookIdNum int
	bookIdNum, _ = strconv.Atoi(bookId)
	err := h.services.BookAct.DeleteBook(c.Request().Context(), bookIdNum)
	if err != nil {
		log.WithFields(log.Fields{
			"Error get book": err,
			"book ID":        bookId,
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
