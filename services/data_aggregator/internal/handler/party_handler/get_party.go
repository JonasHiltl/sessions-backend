package partyhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/party"
	sg "github.com/jonashiltl/sessions-backend/packages/grpc/story"
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/services/data_aggregator/internal/datastruct"
)

func (h partyGatewayHandler) GetParty(c *fiber.Ctx) error {
	id := c.Params("id")

	p, err := h.pc.GetParty(c.Context(), &party.GetPartyRequest{PartyId: id})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	profileRes, _ := h.uc.GetProfile(c.Context(), &ug.GetProfileRequest{Id: p.UserId})

	storyRes, _ := h.sc.GetByParty(c.Context(), &sg.GetByPartyRequest{PartyId: p.Id})

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
