package storyhandler

import (
	"github.com/gofiber/fiber/v2"
	sg "github.com/jonashiltl/sessions-backend/packages/grpc/story"
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/services/data_aggregator/internal/datastruct"
)

func (h storyGatewayHandler) GetStory(c *fiber.Ctx) error {
	sId := c.Params("id")

	s, err := h.sc.GetStory(c.Context(), &sg.GetStoryRequest{StoryId: sId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	// Get all profiles of the tagged people and the story creator in one call
	ids := append(s.TaggedFriends, s.UserId)

	profilesRes, err := h.uc.GetManyProfiles(c.Context(), &ug.GetManyProfilesRequest{Ids: ids})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	// Remove the creator of the story from the returned array and create a filtered list with only the profiles of the tagged people.
	// Separately store the profile of the creator of the story
	var profile *ug.Profile
	var taggedFriends []*ug.Profile
	for _, p := range profilesRes.Profiles {
		if p.Id != s.UserId {
			taggedFriends = append(taggedFriends, p)
		} else {
			profile = p
		}
	}

	res := datastruct.AggregatedStory{
		Id:            s.Id,
		PartyId:       s.PartyId,
		Creator:       profile,
		Lat:           s.Lat,
		Long:          s.Long,
		Url:           s.Url,
		TaggedFriends: taggedFriends,
		CreatedAt:     s.CreatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
