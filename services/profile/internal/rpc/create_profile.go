package handler

import (
	"context"

	"github.com/jonashiltl/sessions-backend/packages/events"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/services/profile/internal/dto"
)

func (s *profileServer) CreateProfile(c context.Context, req *pg.CreateProfileRequest) (*pg.Profile, error) {
	du := dto.Profile{
		Id:        req.Id,
		Username:  req.Username,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Avatar:    req.Avatar,
	}

	p, err := s.ps.Create(c, du)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	s.stream.PublishEvent(events.ProfileCreated{
		Id:        p.Id.Hex(),
		Username:  p.Username,
		Firstname: p.Firstname,
		Lastname:  p.Lastname,
		Avatar:    p.Avatar,
	})

	return p.ToGRPCProfile(), nil
}
