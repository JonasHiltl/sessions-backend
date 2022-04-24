package commenthandler

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	cg "github.com/jonashiltl/sessions-backend/packages/grpc/comment"
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/services/data_aggregator/internal/datastruct"
)

func (h commentGatewayHandler) GetReplyByComment(c *fiber.Ctx) error {
	cId := c.Params("id")
	nextPage := c.Query("nextPage")

	limitStr := c.Query("limit")
	limit, _ := strconv.ParseUint(limitStr, 10, 32)

	rs, err := h.cc.GetReplyByComment(c.Context(), &cg.GetReplyByCommentRequest{CommentId: cId, NextPage: nextPage, Limit: uint32(limit)})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	var replyAuthors []string
	for _, c := range rs.Replies {
		replyAuthors = append(replyAuthors, c.AuthorId)
	}

	ps, err := h.uc.GetManyProfilesMap(c.Context(), &ug.GetManyProfilesMapRequest{Ids: utils.UniqueStringSlice(replyAuthors)})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	log.Printf("%v+", ps.Profiles)

	aggR := make([]datastruct.AggregatedReply, len(rs.Replies))
	for _, r := range rs.Replies {
		aggR = append(aggR, datastruct.AggregatedReply{
			Id:        r.Id,
			CommentId: r.CommentId,
			Author:    ps.Profiles[r.AuthorId],
			Body:      r.Body,
			CreatedAt: r.CreatedAt,
		})
	}

	res := datastruct.PagedAggregatedReply{
		Replies:  aggR,
		NextPage: rs.NextPage,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
