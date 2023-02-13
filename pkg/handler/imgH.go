package handler

import (
	"EFpractic2/models"
	"fmt"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
)

// CreateImg used to create img
func (h *Handler) CreateImg(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}

	dst, err := os.Create("pkg/public/assets/" + file.Filename)
	if err != nil {
		return err
	}

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	img := models.Image{}
	img.ImageRoute = "pkg/public/assets/" + file.Filename
	err = h.serviceImg.CreateImg(c.Request().Context(), &img)
	if err != nil {
		log.WithFields(log.Fields{
			"Error create img": err,
			"img":              img,
		}).Info("CREATE img request")
		return echo.NewHTTPError(http.StatusBadRequest, "img creating failed")
	}
	defer src.Close()
	defer dst.Close()
	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully</p>", file.Filename))
}

// GetImages used to get all image's routes
func (h *Handler) GetImages(c echo.Context) error {
	images, err := h.serviceImg.GetImages(c.Request().Context())
	if err != nil {
		log.WithFields(log.Fields{
			"Error get all images": err,
			"images":               images,
		}).Info("GET ALL images request")
		return echo.NewHTTPError(http.StatusForbidden, "error while getting images")
	}
	return c.JSON(http.StatusOK, images)
}

func (h *Handler) GetCurrentImage(c echo.Context) error {
	var image models.Image
	err := c.Bind(&image)
	if err != nil {
		log.WithFields(log.Fields{
			"Error get image route": err,
		}).Info("get image route request")
		return echo.NewHTTPError(http.StatusForbidden, "error while getting single image")
	}
	return c.Attachment(image.ImageRoute, image.ImageID.String())
}
