package authhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/auth"
	"github.com/jonashiltl/sessions-backend/packages/grpc/profile"
)

type RegisterRequest struct {
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	Username  string `json:"username,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
}

type RegisterResponse struct {
	Token   string           `json:"token,omitempty"`
	Profile *profile.Profile `json:"profile,omitempty"`
}

func (h *authGatewayHandler) Register(c *fiber.Ctx) error {
	req := new(RegisterRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	u, err := h.authClient.Register(c.Context(), &auth.RegisterRequest{Email: req.Email, Password: req.Password})
	if err != nil {
		return err
	}

	profileRes, err := h.profileClient.CreateProfile(c.Context(), &profile.CreateProfileRequest{Id: u.AuthUser.Id, Username: req.Username, Firstname: req.Firstname, Lastname: req.Lastname})
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(RegisterResponse{Token: u.Token, Profile: profileRes})
}
