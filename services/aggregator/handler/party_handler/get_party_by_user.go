package partyhandler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/party"
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/services/aggregator/datastruct"
)

func (h partyGatewayHandler) GetPartyByUser(c *fiber.Ctx) error {
	uId := c.Params("id")
	nextPage := c.Query("nextPage")

	limitStr := c.Query("limit")
	limit, _ := strconv.ParseUint(limitStr, 10, 32)

	partyRes, err := h.pc.GetByUser(c.Context(), &party.GetByUserRequest{UserId: uId, NextPage: nextPage, Limit: uint32(limit)})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	// get the profile of the party creator
	profilesRes, _ := h.uc.GetProfile(c.Context(), &ug.GetProfileRequest{Id: uId})

	aggP := make([]datastruct.AggregatedParty, len(partyRes.Parties))
	for i, p := range partyRes.Parties {
		aggP[i] = datastruct.AggregatedParty{
			Id:            p.Id,
			Creator:       profilesRes,
			Title:         p.Title,
			IsPublic:      p.IsPublic,
			Lat:           p.Lat,
			Long:          p.Long,
			StreetAddress: p.StreetAddress,
			PostalCode:    p.PostalCode,
			State:         p.State,
			Country:       p.Country,
			// TODO: we might want to fetch some stories of the party but would have to do this for all party returned of this user
			// Stories:
			StartDate: p.StartDate,
			CreatedAt: p.CreatedAt,
		}
	}

	res := datastruct.PagedAggregatedParty{
		Parties:  aggP,
		NextPage: partyRes.NextPage,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
