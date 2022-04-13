package rpc

import (
	"context"

	common "github.com/jonashiltl/sessions-backend/packages/grpc/common"
	sg "github.com/jonashiltl/sessions-backend/packages/grpc/story"
	"github.com/jonashiltl/sessions-backend/packages/utils"
)

func (s *storyServer) DeleteStory(c context.Context, req *sg.DeleteStoryRequest) (*common.MessageResponse, error) {
	err := s.sService.Delete(c, req.RequesterId, req.SId)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &common.MessageResponse{Message: "Story removed"}, nil
}
