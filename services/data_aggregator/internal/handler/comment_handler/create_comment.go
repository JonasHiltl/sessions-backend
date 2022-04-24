package commenthandler

import (
	"github.com/gofiber/fiber/v2"
	cg "github.com/jonashiltl/sessions-backend/packages/grpc/comment"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/packages/utils/middleware"
)

func (h commentGatewayHandler) CreateComment(ctx *fiber.Ctx) error {
	req := new(cg.CreateCommentRequest)
	if err := ctx.BodyParser(req); err != nil {
		return err
	}

	pId := ctx.Params("id")
	user := middleware.ParseUser(ctx)

	req.AuthorId = user.Sub
	req.PartyId = pId

	c, err := h.cc.CreateComment(ctx.Context(), req)
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return ctx.Status(fiber.StatusCreated).JSON(c)
}
