package rpc

import (
	"context"

	sg "github.com/jonashiltl/sessions-backend/packages/grpc/story"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *storyServer) GetStory(c context.Context, req *sg.GetStoryRequest) (*sg.PublicStory, error) {
	if req.SId == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid Story id")
	}

	story, err := s.sService.Get(c, req.SId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Story not found")
	}

	return story.ToPublicStory(), nil
}
