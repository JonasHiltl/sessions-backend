package handler

import (
	"context"

	"github.com/jonashiltl/sessions-backend/packages/comutils"
	sg "github.com/jonashiltl/sessions-backend/packages/grpc/story"
)

func (s *storyServer) PresignURL(c context.Context, req *sg.PresignURLRequest) (*sg.PresignURLResponse, error) {

	url, err := s.us.PresignURL(c, req.Key)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	return &sg.PresignURLResponse{Url: url}, nil
}
