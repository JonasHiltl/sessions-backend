package partyhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/comutils"
	"github.com/jonashiltl/sessions-backend/packages/grpc/party"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	sg "github.com/jonashiltl/sessions-backend/packages/grpc/story"
	"github.com/jonashiltl/sessions-backend/services/data_aggregator/internal/datastruct"
)

func (h *partyGatewayHandler) GetParty(c *fiber.Ctx) error {
	id := c.Params("id")

	p, err := h.partyClient.GetParty(c.Context(), &party.GetPartyRequest{PartyId: id})
	if err != nil {
		return comutils.ToHTTPError(err)
	}

	profileRes, err := h.profileClient.GetProfile(c.Context(), &pg.GetProfileRequest{Id: p.UserId})
	if err != nil {
		return comutils.ToHTTPError(err)
	}

	storyRes, err := h.storyClient.GetByParty(c.Context(), &sg.GetByPartyRequest{PartyId: p.Id})
	if err != nil {
		return comutils.ToHTTPError(err)
	}

	res := datastruct.AggregatedParty{
		Id:            p.Id,
		Creator:       profileRes,
		Title:         p.Title,
		IsPublic:      p.IsPublic,
		Lat:           p.Lat,
		Long:          p.Long,
		StreetAddress: p.StreetAddress,
		PostalCode:    p.PostalCode,
		State:         p.State,
		Country:       p.Country,
		Stories:       storyRes.Stories,
		StartDate:     p.StartDate,
		CreatedAt:     p.CreatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
