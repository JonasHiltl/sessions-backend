package handler

import (
	"context"

	sg "github.com/jonashiltl/sessions-backend/packages/grpc/story"
)

func (s *storyServer) PresignURL(c context.Context, req *sg.PresignURLRequest) (*sg.PresignURLResponse, error) {

	url, err := s.us.PresignURL(c, req.Key)
	if err != nil {
		return nil, err
	}

	return &sg.PresignURLResponse{Url: url}, nil
}
