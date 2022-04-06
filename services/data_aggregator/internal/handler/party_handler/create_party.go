package partyhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/party"
)

func (h *partyGatewayHandler) CreateParty(c *fiber.Ctx) error {
	req := new(party.CreatePartyRequest)

	res, err := h.c.CreateParty(c.Context(), req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
