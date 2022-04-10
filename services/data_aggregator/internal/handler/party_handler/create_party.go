package partyhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/comutils"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
	"github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	sg "github.com/jonashiltl/sessions-backend/packages/grpc/story"
	"github.com/jonashiltl/sessions-backend/services/data_aggregator/internal/datastruct"
)

func (h *partyGatewayHandler) CreateParty(c *fiber.Ctx) error {
	req := new(pg.CreatePartyRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	p, err := h.partyClient.CreateParty(c.Context(), req)
	if err != nil {
		return comutils.ToHTTPError(err)
	}

	profileRes, err := h.profileClient.GetProfile(c.Context(), &profile.GetProfileRequest{Id: p.UserId})
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
