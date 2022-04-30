package commenthandler

import (
	"github.com/gofiber/fiber/v2"
	cg "github.com/jonashiltl/sessions-backend/packages/grpc/comment"
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/packages/utils/middleware"
	"github.com/jonashiltl/sessions-backend/services/data_aggregator/internal/datastruct"
)

func (h commentGatewayHandler) CreateComment(c *fiber.Ctx) error {
	req := new(cg.CreateCommentRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	pId := c.Params("id")
	user := middleware.ParseUser(c)

	req.AuthorId = user.Sub
	req.PartyId = pId

	co, err := h.cc.CreateComment(c.Context(), req)
	if err != nil {
		return utils.ToHTTPError(err)
	}

	profileRes, _ := h.uc.GetProfile(c.Context(), &ug.GetProfileRequest{Id: co.AuthorId})

	ac := datastruct.AggregatedComment{
		Id:        co.Id,
		PartyId:   co.PartyId,
		Author:    profileRes,
		Body:      co.Body,
		CreatedAt: co.CreatedAt,
	}

	return c.Status(fiber.StatusCreated).JSON(ac)
}
