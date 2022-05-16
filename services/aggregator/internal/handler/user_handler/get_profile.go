package profilehandler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	rg "github.com/jonashiltl/sessions-backend/packages/grpc/relation"
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/packages/utils/middleware"
	"github.com/jonashiltl/sessions-backend/services/aggregator/internal/datastruct"
)

func (h userGatewayHandler) GetProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	user := middleware.ParseUser(c)

	profile, err := h.uc.GetProfile(c.Context(), &ug.GetProfileRequest{Id: id})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	if profile == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Profile not found")
	}

	log.Println(id)
	log.Println(user.Sub)

	var relation *rg.FriendRelation
	// if somebody wants the Profile of somebody else we also return the friendship status between them two
	if id != user.Sub {
		relation, _ = h.rc.GetFriendRelation(c.Context(), &rg.GetFriendRelationRequest{UserId: user.Sub, FriendId: id})
	}

	res := datastruct.AggregatedProfile{
		Id:          profile.Id,
		Username:    profile.Username,
		Firstname:   profile.Firstname,
		Lastname:    profile.Lastname,
		Avatar:      profile.Avatar,
		FriendCount: profile.FriendCount,
	}

	log.Println(relation)

	if relation != nil {
		st := datastruct.FriendshipStatus{}

		if relation.Accepted {
			st.IsFriend = true
			st.AcceptedAt = relation.AcceptedAt
		} else {
			st.OutgoingRequest = true
			st.RequestedAt = relation.RequestedAt
		}

		res.FriendshipStatus = st
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
