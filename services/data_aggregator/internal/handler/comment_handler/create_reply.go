package commenthandler

import (
	"github.com/gofiber/fiber/v2"
	cg "github.com/jonashiltl/sessions-backend/packages/grpc/comment"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/packages/utils/middleware"
)

func (h commentGatewayHandler) CreateReply(c *fiber.Ctx) error {
	req := new(cg.CreateReplyRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	cId := c.Params("id")
	user := middleware.ParseUser(c)

	req.AuthorId = user.Sub
	req.CommentId = cId

	res, err := h.cc.CreateReply(c.Context(), req)
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}
