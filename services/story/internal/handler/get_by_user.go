package handler

import (
	"context"
	"encoding/base64"

	"github.com/jonashiltl/sessions-backend/packages/comutils"
	sg "github.com/jonashiltl/sessions-backend/packages/grpc/story"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *storyServer) GetByUser(c context.Context, req *sg.GetByUserRequest) (*sg.PagedStories, error) {
	p, err := base64.URLEncoding.DecodeString(req.NextPage)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Next Page Param")
	}
	stories, p, err := s.sService.GetByUser(c, req.UId, p)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	nextPage := base64.URLEncoding.EncodeToString(p)

	var result []*sg.PublicStory
	for _, s := range stories {
		result = append(result, s.ToPublicStory())
	}

	return &sg.PagedStories{Stories: result, NextPage: nextPage}, nil
}
