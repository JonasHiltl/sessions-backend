package handler

import (
	"log"
	"net/http"

	"github.com/jonashiltl/sessions-backend/services/party/internal/datastruct"
	"github.com/labstack/echo/v4"
)

// @Summary Get parties of a user
// @Description Returns a list of parties of a user
// @Produce json
// @Param uId path string true "User id"
// @Success 200 {array} datastruct.PublicParty
// @Failure 400 {object} echo.HTTPError
// @Router /user/{uId} [get]
func (a *httpApp) GetByUser(c echo.Context) error {
	uId := c.Param("uId")

	var page []byte

	ps, nextPage, err := a.partyService.GetByUser(c.Request().Context(), uId, page)
	log.Println("Next Page: ", nextPage)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var pp []datastruct.PublicParty
	for _, p := range ps {
		pp = append(pp, p.ToPublicParty())
	}

	return c.JSON(http.StatusOK, pp)
}
