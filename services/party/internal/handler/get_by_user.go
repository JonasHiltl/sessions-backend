package handler

import (
	"encoding/base64"
	"log"
	"net/http"

	"github.com/jonashiltl/sessions-backend/services/party/internal/datastruct"
	"github.com/labstack/echo/v4"
)

// @Summary Get parties of a user
// @Description Returns a list of parties of a user
// @Produce json
// @Param uId path string true "User id"
// @Success 200 {object} datastruct.PagedParties
// @Failure 400 {object} echo.HTTPError
// @Router /user/{uId} [get]
func (a *httpApp) GetByUser(c echo.Context) error {
	uId := c.Param("uId")
	pageQuery := c.QueryParam("nextPage")

	p, err := base64.URLEncoding.DecodeString(pageQuery)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Next Page Param")
	}

	ps, nextPage, err := a.partyService.GetByUser(c.Request().Context(), uId, p)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	str := base64.URLEncoding.EncodeToString(nextPage)

	var pp []datastruct.PublicParty
	for _, p := range ps {
		pp = append(pp, p.ToPublicParty())
	}

	return c.JSON(http.StatusOK, datastruct.PagedParties{Parties: pp, NextPage: str})
}
