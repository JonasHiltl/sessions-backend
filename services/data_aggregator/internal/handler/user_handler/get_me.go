package profilehandler

import (
	"github.com/gofiber/fiber/v2"
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/utils"

	"github.com/jonashiltl/sessions-backend/packages/utils/middleware"
)

func (h userGatewayHandler) GetMe(c *fiber.Ctx) error {
	user := middleware.ParseUser(c)

	res, err := h.uc.GetMe(c.Context(), &ug.GetMeRequest{Id: user.Sub})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
