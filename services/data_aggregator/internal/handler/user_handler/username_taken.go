package profilehandler

import (
	"github.com/gofiber/fiber/v2"
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/utils"
)

type UsernameTakenResponse struct {
	Taken bool `json:"taken"`
}

func (h userGatewayHandler) UsernameTaken(c *fiber.Ctx) error {
	uName := c.Params("username")

	res, err := h.uc.UsernameTaken(c.Context(), &ug.UsernameTakenRequest{Username: uName})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(&UsernameTakenResponse{Taken: res.Taken})
}
