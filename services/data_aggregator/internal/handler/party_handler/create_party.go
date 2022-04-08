package partyhandler

import (
	"github.com/gofiber/fiber/v2"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
	"github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	sg "github.com/jonashiltl/sessions-backend/packages/grpc/story"
	"github.com/jonashiltl/sessions-backend/services/data_aggregator/internal/datastruct"
)

func (h *partyGatewayHandler) CreateParty(c *fiber.Ctx) error {
	req := new(pg.CreatePartyRequest)

	p, err := h.partyClient.CreateParty(c.Context(), req)
	if err != nil {
		return err
	}

	profileRes, err := h.profileClient.GetProfile(c.Context(), &profile.GetProfileRequest{Id: p.UserId})
	if err != nil {
		return err
	}

	storyRes, err := h.storyClient.GetByParty(c.Context(), &sg.GetByPartyRequest{PartyId: p.Id})
	if err != nil {
		return err
	}

	res := datastruct.AggregatedParty{
		Id:        p.Id,
		Creator:   profileRes,
		Title:     p.Title,
		IsPublic:  p.IsPublic,
		Lat:       p.Lat,
		Long:      p.Long,
		Stories:   storyRes.Stories,
		CreatedAt: p.CreatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
