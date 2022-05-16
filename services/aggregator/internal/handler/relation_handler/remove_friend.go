package relationhandler

import (
	"github.com/gofiber/fiber/v2"
	rg "github.com/jonashiltl/sessions-backend/packages/grpc/relation"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/packages/utils/middleware"
)

func (h relationGatewayHandler) RemoveFriend(c *fiber.Ctx) error {
	user := middleware.ParseUser(c)
	uId := c.Params("id")

	if user.Sub == uId {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Friend Id")
	}

	ok, err := h.relationClient.RemoveFriend(c.Context(), &rg.RemoveFriendRequest{UserId: user.Sub, FriendId: uId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(ok)
}
