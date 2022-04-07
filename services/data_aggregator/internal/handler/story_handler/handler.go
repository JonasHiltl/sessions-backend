package storyhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"github.com/jonashiltl/sessions-backend/packages/grpc/story"
)

type storyGatewayHandler struct {
	storyClient   story.StoryServiceClient
	profileClient profile.ProfileServiceClient
}

type StoryGatewayHandler interface {
	CreateStory(c *fiber.Ctx) error
	GetStory(c *fiber.Ctx) error
	GetStoryByParty(c *fiber.Ctx) error
	GetStoryByUser(c *fiber.Ctx) error
	DeleteStory(c *fiber.Ctx) error
	PresignURL(c *fiber.Ctx) error
}

func NewStoryGatewayHandler(storyClient story.StoryServiceClient, profileClient profile.ProfileServiceClient) StoryGatewayHandler {
	return &storyGatewayHandler{
		storyClient:   storyClient,
		profileClient: profileClient,
	}
}
