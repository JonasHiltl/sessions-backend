package authhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/auth"
)

type RegisterRequest struct {
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	Username  string `json:"username,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
}

func (h *authGatewayHandler) Register(c *fiber.Ctx) error {
	req := new(RegisterRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	res, err := h.authClient.Register(c.Context(), &auth.RegisterRequest{Email: req.Email, Password: req.Password})
	if err != nil {
		return err
	}

	// TODO: Also create profile with

	return c.Status(fiber.StatusOK).JSON(res)
}
