package partyhandler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/party"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/services/data_aggregator/internal/datastruct"
)

func (h *partyGatewayHandler) GetPartyByUser(c *fiber.Ctx) error {
	uId := c.Params("id")
	nextPage := c.Query("nextPage")

	limitStr := c.Query("limit")
	limit, err := strconv.ParseUint(limitStr, 10, 32)
	if err != nil {
		limit = 0
	}

	partyRes, err := h.partyClient.GetByUser(c.Context(), &party.GetByUserRequest{UserId: uId, NextPage: nextPage, Limit: uint32(limit)})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	// get the profile of the party creator
	profilesRes, err := h.profileClient.GetProfile(c.Context(), &pg.GetProfileRequest{Id: uId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	var aggParties []datastruct.AggregatedParty
	for _, p := range partyRes.Parties {
		aggParties = append(aggParties, datastruct.AggregatedParty{
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
		})
	}

	res := datastruct.PagedAggregatedParty{
		Parties:  aggParties,
		NextPage: partyRes.NextPage,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
