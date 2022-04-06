package partyhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/party"
)

func (h *partyGatewayHandler) GetPartyByUser(c *fiber.Ctx) error {
	uId := c.Params("id")
	nextPage := c.Query("nextPage")

	res, err := h.c.GetByUser(c.Context(), &party.GetByUserRequest{UId: uId, NextPage: nextPage})
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
