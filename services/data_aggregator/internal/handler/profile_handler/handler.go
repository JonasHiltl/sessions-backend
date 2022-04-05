package profilehandler

import "github.com/jonashiltl/sessions-backend/packages/grpc/profile"

type profileGatewayHandler struct {
	c profile.ProfileServiceClient
}

type ProfileGatewayHandler interface {
}

func NewProfileGatewayHandler(pc profile.ProfileServiceClient) ProfileGatewayHandler {
	return &profileGatewayHandler{
		c: pc,
	}
}
