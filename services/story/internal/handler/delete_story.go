package handler

import (
	"context"

	"github.com/jonashiltl/sessions-backend/packages/comutils"
	"github.com/jonashiltl/sessions-backend/packages/comutils/middleware"
	common "github.com/jonashiltl/sessions-backend/packages/grpc/common"
	sg "github.com/jonashiltl/sessions-backend/packages/grpc/story"
)

func (s *storyServer) DeleteStory(c context.Context, req *sg.GetStoryRequest) (*common.MessageResponse, error) {
	me, err := middleware.ParseUser(c)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	err = s.sService.Delete(c, me.Sub, req.SId)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	return &common.MessageResponse{Message: "Story removed"}, nil
}
