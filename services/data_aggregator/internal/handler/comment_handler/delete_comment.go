package commenthandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/comment"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/packages/utils/middleware"
)

func (h commentGatewayHandler) DeleteComment(c *fiber.Ctx) error {
	user := middleware.ParseUser(c)

	cId := c.Params("id")
	pId := c.Params("pId")

	res, err := h.cc.DeleteComment(c.Context(), &comment.DeleteCommentRequest{AuthorId: user.Sub, PartyId: pId, CommentId: cId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)

}
