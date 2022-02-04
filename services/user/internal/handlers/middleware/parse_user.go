package middleware

import (
	"encoding/base64"
	"encoding/json"
	"errors"

	"github.com/jonashiltl/sessions-backend/services/user/internal/datastruct"
	"github.com/labstack/echo/v4"
)

func ParseUser(c echo.Context) (*datastruct.JwtPayload, error) {
	p := c.Request().Header.Get("jwt_payload")
	if p == "" {
		return nil, errors.New("not logged in")
	}

	stringBytes, err := base64.StdEncoding.DecodeString(p)
	if err != nil {
		return nil, err
	}

	result := datastruct.JwtPayload{}
	json.Unmarshal(stringBytes, &result)

	return &result, nil

}
