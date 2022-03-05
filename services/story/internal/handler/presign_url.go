package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type presignRes struct {
	url string
}

// @Summary Get a presigned S3 URL
// @Description Returns a url to update a video story to
// @Accept json
// @Produce json
// @Param key path string true "Key of file to upload"
// @Success 200 {object} presignRes
// @Failure 400 {object} echo.HTTPError
// @Router /presign/{key} [get]
func (a *httpApp) PresignURL(c echo.Context) error {
	key := c.Param("key")

	url, err := a.us.PresignURL(c.Request().Context(), key)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, presignRes{url: url})
}
