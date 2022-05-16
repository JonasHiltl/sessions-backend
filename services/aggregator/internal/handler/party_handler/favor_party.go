package partyhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/party"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/packages/utils/middleware"
)

func (h partyGatewayHandler) FavorParty(c *fiber.Ctx) error {
	user := middleware.ParseUser(c)

	pId := c.Params("id")

	f, err := h.pc.FavorParty(c.Context(), &party.FavorPartyRequest{UserId: user.Sub, PartyId: pId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(f)
}
