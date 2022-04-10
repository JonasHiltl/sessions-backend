package partyhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/comutils"
	"github.com/jonashiltl/sessions-backend/packages/grpc/party"
)

func (h *partyGatewayHandler) GetPartyByUser(c *fiber.Ctx) error {
	uId := c.Params("id")
	nextPage := c.Query("nextPage")

	res, err := h.partyClient.GetByUser(c.Context(), &party.GetByUserRequest{UserId: uId, NextPage: nextPage})
	if err != nil {
		return comutils.ToHTTPError(c, err)
	}

	// TODO: Decide if it's necessary to always get one/multiple stories of a party to have anything to show in the Frontend when viewing a Party

	return c.Status(fiber.StatusOK).JSON(res)
}
