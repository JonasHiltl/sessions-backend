package partyhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/party"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/packages/utils/middleware"
)

func (h partyGatewayHandler) DefavorParty(c *fiber.Ctx) error {
	user := middleware.ParseUser(c)

	pId := c.Params("id")

	ok, err := h.pc.DefavorParty(c.Context(), &party.FavorPartyRequest{UserId: user.Sub, PartyId: pId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(ok)
}
