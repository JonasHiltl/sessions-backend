package storyhandler

import (
	"github.com/gofiber/fiber/v2"
	sc "github.com/jonashiltl/sessions-backend/packages/grpc/story"
	uc "github.com/jonashiltl/sessions-backend/packages/grpc/user"
)

type storyGatewayHandler struct {
	sc sc.StoryServiceClient
	uc uc.UserServiceClient
}

type StoryGatewayHandler interface {
	CreateStory(c *fiber.Ctx) error
	GetStory(c *fiber.Ctx) error
	GetStoryByParty(c *fiber.Ctx) error
	GetStoryByUser(c *fiber.Ctx) error
	DeleteStory(c *fiber.Ctx) error
	PresignURL(c *fiber.Ctx) error
}

func NewStoryGatewayHandler(sc sc.StoryServiceClient, uc uc.UserServiceClient) StoryGatewayHandler {
	return &storyGatewayHandler{
		sc: sc,
		uc: uc,
	}
}
