package rpc

import (
	"context"
	"encoding/base64"

	sg "github.com/jonashiltl/sessions-backend/packages/grpc/story"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s storyServer) GetByParty(c context.Context, req *sg.GetByPartyRequest) (*sg.PagedStories, error) {
	p, err := base64.URLEncoding.DecodeString(req.NextPage)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Next Page Param")
	}

	stories, p, err := s.sService.GetByParty(c, req.PartyId, p, req.Limit)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	nextPage := base64.URLEncoding.EncodeToString(p)

	var result []*sg.PublicStory
	for _, s := range stories {
		result = append(result, s.ToPublicStory())
	}

	return &sg.PagedStories{Stories: result, NextPage: nextPage}, nil
}
