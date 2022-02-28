package handler

import (
	"net/http"

	"github.com/jonashiltl/sessions-backend/packages/comutils/middleware"
	"github.com/jonashiltl/sessions-backend/services/story/internal/dto"
	"github.com/labstack/echo/v4"
)

type StoryCreateRequest struct {
	PId           string   `json:"p_id,omitempty"           validate:"required"`
	Url           string   `json:"url,omitempty"            validate:"required"`
	Lat           float32  `json:"lat,omitempty"            validate:"required,latitude"`
	Long          float32  `json:"long,omitempty"           validate:"required,longitude"`
	TaggedFriends []string `json:"tagged_friends,omitempty"`
}

func (a *httpApp) CreateStory(c echo.Context) error {
	var req StoryCreateRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Couldn't find request body")
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	me, err := middleware.ParseUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	d := dto.Story{
		PId:           req.PId,
		UId:           me.Sub,
		Lat:           float64(req.Lat),
		Long:          float64(req.Long),
		Url:           req.Url,
		TaggedFriends: req.TaggedFriends,
	}

	story, err := a.sService.Create(c.Request().Context(), d)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, story.ToPublicStory())
}
