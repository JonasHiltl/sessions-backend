package relationhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/relation"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/packages/utils/middleware"
)

func (h *relationGatewayHandler) AcceptFriend(c *fiber.Ctx) error {
	user := middleware.ParseUser(c)

	uId := c.Params("id")

	if user.Sub == uId {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Friend Id")
	}

	fr, err := h.relationClient.AcceptFriend(c.Context(), &relation.AcceptFriendRequest{UserId: uId, FriendId: user.Sub})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(fr)
}
