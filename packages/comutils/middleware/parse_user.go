package middleware

import (
	"encoding/base64"
	"encoding/json"
	"errors"

	"github.com/jonashiltl/sessions-backend/packages/comtypes"
	"github.com/labstack/echo/v4"
)

func ParseUser(c echo.Context) (comtypes.JwtPayload, error) {
	p := c.Request().Header.Get("jwt_payload")
	if p == "" {
		return comtypes.JwtPayload{}, errors.New("not logged in")
	}

	stringBytes, err := base64.StdEncoding.DecodeString(p)
	if err != nil {
		return comtypes.JwtPayload{}, err
	}

	result := comtypes.JwtPayload{}
	json.Unmarshal(stringBytes, &result)

	return result, nil

}
