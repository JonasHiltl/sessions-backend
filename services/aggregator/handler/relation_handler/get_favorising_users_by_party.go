package relationhandler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	rg "github.com/jonashiltl/sessions-backend/packages/grpc/relation"
	"github.com/jonashiltl/sessions-backend/packages/grpc/user"

	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/services/aggregator/datastruct"
)

func (h relationGatewayHandler) GetFavorisingUsersByParty(c *fiber.Ctx) error {
	pId := c.Params("id")
	nextPage := c.Query("nextPage")

	limitStr := c.Query("limit")
	limit, _ := strconv.ParseUint(limitStr, 10, 32)

	fpRes, err := h.rc.GetFavorisingUsersByParty(c.Context(), &rg.GetFavorisingUsersByPartyRequest{PartyId: pId, NextPage: nextPage, Limit: uint32(limit)})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	var userIds []string
	for _, fp := range fpRes.FavoriteParties {
		userIds = append(userIds, fp.UserId)
	}

	pRes, _ := h.uc.GetManyProfilesMap(c.Context(), &user.GetManyProfilesRequest{Ids: userIds})

	aggFP := make([]datastruct.AggregatedFavorisingUsers, len(fpRes.FavoriteParties))
	for i, fp := range fpRes.FavoriteParties {
		aggFP[i] = datastruct.AggregatedFavorisingUsers{
			User:        pRes.Profiles[fp.UserId],
			PartyId:     fp.PartyId,
			FavoritedAt: fp.FavoritedAt,
		}
	}

	res := datastruct.PagedAggregatedFavorisingUsers{
		FavoriteParties: aggFP,
		NextPage:        fpRes.NextPage,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
