package partyhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/party"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/packages/utils/middleware"
)

func (h partyGatewayHandler) DeleteParty(c *fiber.Ctx) error {
	user := middleware.ParseUser(c)

	pId := c.Params("id")

	res, err := h.pc.DeleteParty(c.Context(), &party.DeletePartyRequest{RequesterId: user.Sub, PartyId: pId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
