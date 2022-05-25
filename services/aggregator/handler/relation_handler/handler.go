package relationhandler

import (
	"github.com/gofiber/fiber/v2"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
	rg "github.com/jonashiltl/sessions-backend/packages/grpc/relation"
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
)

type relationGatewayHandler struct {
	rc rg.RelationServiceClient
	pc pg.PartyServiceClient
	uc ug.UserServiceClient
}

type RelationGatewayHandler interface {
	FriendRequest(c *fiber.Ctx) error
	AcceptFriend(c *fiber.Ctx) error
	RemoveFriend(c *fiber.Ctx) error

	FavorParty(c *fiber.Ctx) error
	DefavorParty(c *fiber.Ctx) error
	GetFavoritePartiesByUser(c *fiber.Ctx) error
	GetFavorisingUsersByParty(c *fiber.Ctx) error
}

func NewRelationGatewayHandler(rc rg.RelationServiceClient, pc pg.PartyServiceClient, uc ug.UserServiceClient) RelationGatewayHandler {
	return &relationGatewayHandler{
		rc: rc,
		pc: pc,
		uc: uc,
	}
}
