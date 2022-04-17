package authhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/auth"
	"github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"github.com/jonashiltl/sessions-backend/packages/stream"
)

type authGatewayHandler struct {
	authClient    auth.AuthServiceClient
	profileClient profile.ProfileServiceClient
	stream        stream.Stream
}

type AuthGatewayHandler interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	GoogleLogin(c *fiber.Ctx) error
	VerifyEmail(c *fiber.Ctx) error
}

func NewAuthGatewayHandler(authClient auth.AuthServiceClient, profileClient profile.ProfileServiceClient, stream stream.Stream) AuthGatewayHandler {
	return &authGatewayHandler{
		authClient:    authClient,
		profileClient: profileClient,
		stream:        stream,
	}
}
