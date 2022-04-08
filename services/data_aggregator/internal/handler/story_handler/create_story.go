package storyhandler

import (
	"github.com/gofiber/fiber/v2"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"github.com/jonashiltl/sessions-backend/packages/grpc/story"
	"github.com/jonashiltl/sessions-backend/services/data_aggregator/internal/datastruct"
)

func (h *storyGatewayHandler) CreateStory(c *fiber.Ctx) error {
	req := new(story.CreateStoryRequest)

	s, err := h.storyClient.CreateStory(c.Context(), req)
	if err != nil {
		return err
	}

	// Get all profiles of the tagged people and the story creator in one call
	ids := append(s.TaggedFriends, s.UserId)

	profilesRes, err := h.profileClient.GetManyProfiles(c.Context(), &pg.GetManyProfilesRequest{Ids: ids})
	if err != nil {
		return err
	}

	// Remove the creator of the story from the returned array and create a filtered list with only the profiles of the tagged people.
	// Separately store the profile of the creator of the story
	var profile *pg.Profile
	var taggedFriends []*pg.Profile
	for _, p := range profilesRes.Profiles {
		if p.Id != s.UserId {
			taggedFriends = append(taggedFriends, p)
		} else {
			profile = p
		}
	}

	res := datastruct.StoryAggregated{
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
