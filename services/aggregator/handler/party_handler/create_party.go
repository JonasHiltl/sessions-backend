package partyhandler

import (
	"github.com/gofiber/fiber/v2"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/packages/utils/middleware"
	"github.com/jonashiltl/sessions-backend/services/aggregator/datastruct"
)

func (h partyGatewayHandler) CreateParty(c *fiber.Ctx) error {
	req := new(pg.CreatePartyRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	user := middleware.ParseUser(c)
	req.RequesterId = user.Sub

	p, err := h.pc.CreateParty(c.Context(), req)
	if err != nil {
		return utils.ToHTTPError(err)
	}

	profileRes, _ := h.uc.GetProfile(c.Context(), &ug.GetProfileRequest{Id: p.UserId})

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
		StartDate:     p.StartDate,
		CreatedAt:     p.CreatedAt,
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}
