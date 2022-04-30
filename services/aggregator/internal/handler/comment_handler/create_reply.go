package commenthandler

import (
	"github.com/gofiber/fiber/v2"
	cg "github.com/jonashiltl/sessions-backend/packages/grpc/comment"
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/packages/utils/middleware"
	"github.com/jonashiltl/sessions-backend/services/aggregator/internal/datastruct"
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

	r, err := h.cc.CreateReply(c.Context(), req)
	if err != nil {
		return utils.ToHTTPError(err)
	}

	profileRes, _ := h.uc.GetProfile(c.Context(), &ug.GetProfileRequest{Id: r.AuthorId})

	ar := datastruct.AggregatedReply{
		Id:        r.Id,
		CommentId: r.CommentId,
		Author:    profileRes,
		Body:      r.Body,
		CreatedAt: r.CreatedAt,
	}

	return c.Status(fiber.StatusCreated).JSON(ar)
}
