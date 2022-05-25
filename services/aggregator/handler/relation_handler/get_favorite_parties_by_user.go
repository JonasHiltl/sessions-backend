package relationhandler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/party"
	rg "github.com/jonashiltl/sessions-backend/packages/grpc/relation"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/services/aggregator/datastruct"
)

func (h *relationGatewayHandler) GetFavoritePartiesByUser(c *fiber.Ctx) error {
	uId := c.Params("id")
	nextPage := c.Query("nextPage")

	limitStr := c.Query("limit")
	limit, _ := strconv.ParseUint(limitStr, 10, 32)

	fpRes, err := h.rc.GetFavoritePartiesByUser(c.Context(), &rg.GetFavoritePartiesByUserRequest{UserId: uId, NextPage: nextPage, Limit: uint32(limit)})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	var partyIds []string
	for _, fp := range fpRes.FavoriteParties {
		partyIds = append(partyIds, fp.PartyId)
	}

	pRes, _ := h.pc.GetManyPartiesMap(c.Context(), &party.GetManyPartiesRequest{Ids: partyIds})

	aggFP := make([]datastruct.AggregatedFavoriteParty, len(fpRes.FavoriteParties))
	for i, fp := range fpRes.FavoriteParties {
		aggFP[i] = datastruct.AggregatedFavoriteParty{
			UserId:      fp.UserId,
			Party:       pRes.Parties[fp.PartyId],
			FavoritedAt: fp.FavoritedAt,
		}
	}

	res := datastruct.PagedAggregatedFavoriteParty{
		FavoriteParties: aggFP,
		NextPage:        fpRes.NextPage,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
