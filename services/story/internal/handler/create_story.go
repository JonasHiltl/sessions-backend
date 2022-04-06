package handler

import (
	"context"

	"github.com/jonashiltl/sessions-backend/packages/comutils"
	sg "github.com/jonashiltl/sessions-backend/packages/grpc/story"
	"github.com/jonashiltl/sessions-backend/services/story/internal/dto"
)

type StoryCreateRequest struct {
	PId           string   `json:"p_id,omitempty"           validate:"required"`
	Url           string   `json:"url,omitempty"            validate:"required"`
	Lat           float32  `json:"lat,omitempty"            validate:"required,latitude"`
	Long          float32  `json:"long,omitempty"           validate:"required,longitude"`
	TaggedFriends []string `json:"tagged_friends,omitempty"`
}

func (s *storyServer) CreateStory(c context.Context, req *sg.CreateStoryRequest) (*sg.PublicStory, error) {
	d := dto.Story{
		PId:           req.PId,
		UId:           req.RequesterId,
		Lat:           float64(req.GetLat()),
		Long:          float64(req.GetLong()),
		Url:           req.Url,
		TaggedFriends: req.TaggedFriends,
	}

	story, err := s.sService.Create(c, d)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	return story.ToPublicStory(), err
}
