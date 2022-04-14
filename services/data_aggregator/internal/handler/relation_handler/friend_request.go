package relationhandler

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/relation"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/packages/utils/middleware"
)

func (h *relationGatewayHandler) FriendRequest(c *fiber.Ctx) error {
	user := middleware.ParseUser(c)

	fId := c.Params("id")
	log.Println(strings.TrimSpace(fId))

	if user.Sub == fId {
		return fiber.NewError(fiber.StatusBadRequest, "You can't add yourself")
	}

	fr, err := h.relationClient.FriendRequest(c.Context(), &relation.FriendRequestRequest{UserId: user.Sub, FriendId: fId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(fr)
}
