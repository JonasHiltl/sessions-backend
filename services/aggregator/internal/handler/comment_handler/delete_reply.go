package commenthandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/comment"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/packages/utils/middleware"
)

func (h *commentGatewayHandler) DeleteReply(c *fiber.Ctx) error {
	user := middleware.ParseUser(c)

	cId := c.Params("id")
	rId := c.Params("rId")

	res, err := h.cc.DeleteReply(c.Context(), &comment.DeleteReplyRequest{AuthorId: user.Sub, CommentId: cId, ReplyId: rId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
