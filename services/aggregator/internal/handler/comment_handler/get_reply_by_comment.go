package commenthandler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	cg "github.com/jonashiltl/sessions-backend/packages/grpc/comment"
	"github.com/jonashiltl/sessions-backend/packages/grpc/user"
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/services/aggregator/internal/datastruct"
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
	for _, r := range rs.Replies {
		replyAuthors = append(replyAuthors, r.AuthorId)
	}

	ps, _ := h.uc.GetManyProfilesMap(c.Context(), &ug.GetManyProfilesRequest{Ids: utils.UniqueStringSlice(replyAuthors)})

	aggR := make([]datastruct.AggregatedReply, len(rs.Replies))
	for i, r := range rs.Replies {
		if author, ok := ps.Profiles[r.AuthorId]; ok {
			aggR[i] = datastruct.AggregatedReply{
				Id:        r.Id,
				CommentId: r.CommentId,
				Author:    author,
				Body:      r.Body,
				CreatedAt: r.CreatedAt,
			}
		} else {
			aggR[i] = datastruct.AggregatedReply{
				Id:        r.Id,
				CommentId: r.CommentId,
				Author:    &user.Profile{},
				Body:      r.Body,
				CreatedAt: r.CreatedAt,
			}
		}
	}

	res := datastruct.PagedAggregatedReply{
		Replies:  aggR,
		NextPage: rs.NextPage,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
