package authhandler

import (
	"github.com/gofiber/fiber/v2"
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/utils"
)

type RegisterRequest struct {
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	Username  string `json:"username,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
}

func (h *authGatewayHandler) Register(c *fiber.Ctx) error {
	req := new(ug.RegisterRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	u, err := h.uc.Register(c.Context(), req)
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(u)
}
