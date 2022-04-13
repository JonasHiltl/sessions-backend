package rpc

import (
	"context"

	sg "github.com/jonashiltl/sessions-backend/packages/grpc/story"
	"github.com/jonashiltl/sessions-backend/packages/utils"
)

func (s *storyServer) PresignURL(c context.Context, req *sg.PresignURLRequest) (*sg.PresignURLResponse, error) {

	url, err := s.us.PresignURL(c, req.Key)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &sg.PresignURLResponse{Url: url}, nil
}
