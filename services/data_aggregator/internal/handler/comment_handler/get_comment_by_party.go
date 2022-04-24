package commenthandler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	cg "github.com/jonashiltl/sessions-backend/packages/grpc/comment"
	"github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/services/data_aggregator/internal/datastruct"
)

func (h commentGatewayHandler) GetCommentByParty(c *fiber.Ctx) error {
	pId := c.Params("id")
	nextPage := c.Query("nextPage")

	limitStr := c.Query("limit")
	limit, _ := strconv.ParseUint(limitStr, 10, 32)

	cs, err := h.cc.GetCommentByParty(c.Context(), &cg.GetByPartyRequest{PartyId: pId, NextPage: nextPage, Limit: uint32(limit)})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	var commentAuthors []string
	for _, c := range cs.Comments {
		commentAuthors = append(commentAuthors, c.AuthorId)
	}

	pRes, err := h.uc.GetManyProfilesMap(c.Context(), &user.GetManyProfilesMapRequest{Ids: utils.UniqueStringSlice(commentAuthors)})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	aggC := make([]datastruct.AggregatedComment, len(cs.Comments))
	for i, c := range cs.Comments {
		aggC[i] = datastruct.AggregatedComment{
			Id:        c.Id,
			PartyId:   c.PartyId,
			Author:    pRes.Profiles[c.AuthorId],
			Body:      c.Body,
			CreatedAt: c.CreatedAt,
		}
	}

	res := datastruct.PagedAggregatedComment{
		Comments: aggC,
		NextPage: cs.NextPage,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
