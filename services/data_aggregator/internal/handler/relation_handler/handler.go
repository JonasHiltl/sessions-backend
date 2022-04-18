package relationhandler

import (
	"github.com/gofiber/fiber/v2"
	rg "github.com/jonashiltl/sessions-backend/packages/grpc/relation"
)

type relationGatewayHandler struct {
	relationClient rg.RelationServiceClient
}

type RelationGatewayHandler interface {
	FriendRequest(c *fiber.Ctx) error
	AcceptFriend(c *fiber.Ctx) error
	RemoveFriend(c *fiber.Ctx) error
}

func NewRelationGatewayHandler(relationClient rg.RelationServiceClient) RelationGatewayHandler {
	return &relationGatewayHandler{
		relationClient: relationClient,
	}
}
