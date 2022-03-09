package handler

import (
	"encoding/base64"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/jonashiltl/sessions-backend/services/party/internal/datastruct"
	"github.com/labstack/echo/v4"
)

type GeoSearchBody struct {
	Lat       float64 `json:"lat"       validate:"required,latitude"`
	Long      float64 `json:"long"      validate:"required,longitude"`
	Precision uint    `json:"precision"`
}

// @Summary Search by location
// @Description Get a list of parties near a location
// @Tags GEO
// @Produce json
// @Param lat query float32 true "Latitude"
// @Param long query float32 true "Longitude"
// @Param precision query uint true "Geohash precision"
// @Success 200 {object} datastruct.PagedParties
// @Failure 400 {object} echo.HTTPError
// @Router /near [get]
func (a *httpApp) GeoSearch(c echo.Context) error {
	lString := c.QueryParam("lat")
	loString := c.QueryParam("long")
	pString := c.QueryParam("precision")

	precision, err := strconv.ParseUint(pString, 10, 8)
	if err != nil {
		return errors.New("invalid precision parameter format")
	}
	long, err := strconv.ParseFloat(loString, 64)
	if err != nil {
		return errors.New("invalid long parameter format")
	}
	lat, err := strconv.ParseFloat(lString, 64)
	if err != nil {
		return errors.New("invalid lat parameter format")
	}

	pageQuery := c.QueryParam("nextPage")

	p, err := base64.URLEncoding.DecodeString(pageQuery)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Next Page Param")
	}

	var reqBody GeoSearchBody
	if err := c.Bind(&reqBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Couldn't find request body")
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}

	ps, nextPage, err := a.partyService.GeoSearch(c.Request().Context(), lat, long, uint(precision), p)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	str := base64.URLEncoding.EncodeToString(nextPage)

	var pp []datastruct.PublicParty

	for _, ps := range ps {
		pp = append(pp, ps.ToPublicParty())
	}

	return c.JSON(http.StatusOK, datastruct.PagedParties{Parties: pp, NextPage: str})
}
