package authhandler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/events"
	ag "github.com/jonashiltl/sessions-backend/packages/grpc/auth"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"github.com/jonashiltl/sessions-backend/packages/utils"
)

type RegisterRequest struct {
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	Username  string `json:"username,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
}

type RegisterResponse struct {
	Token   string      `json:"token,omitempty"`
	Profile *pg.Profile `json:"profile,omitempty"`
}

func (h *authGatewayHandler) Register(c *fiber.Ctx) error {
	req := new(RegisterRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	uTaken, err := h.profileClient.UsernameTaken(c.Context(), &pg.UsernameTakenRequest{Username: req.Username})
	if uTaken.Taken || err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Username already taken")
	}
	eTaken, err := h.authClient.EmailTaken(c.Context(), &ag.EmailTakenRequest{Email: req.Email})
	log.Printf("%v", eTaken)
	if eTaken.Taken || err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Email already taken")
	}

	u, err := h.authClient.Register(c.Context(), &ag.RegisterRequest{Email: req.Email, Password: req.Password})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	profileRes, err := h.profileClient.CreateProfile(c.Context(), &pg.CreateProfileRequest{Id: u.AuthUser.Id, Username: req.Username, Firstname: req.Firstname, Lastname: req.Lastname})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	h.stream.PublishEvent(&events.Registered{
		Id:        u.AuthUser.Id,
		Email:     u.AuthUser.Email,
		Username:  profileRes.Username,
		Firstname: profileRes.Firstname,
		Lastname:  profileRes.Lastname,
	})

	return c.Status(fiber.StatusOK).JSON(RegisterResponse{Token: u.Token, Profile: profileRes})
}
