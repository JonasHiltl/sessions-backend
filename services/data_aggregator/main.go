package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/auth"
	"github.com/jonashiltl/sessions-backend/packages/grpc/party"
	"github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"github.com/jonashiltl/sessions-backend/packages/grpc/story"
	"github.com/jonashiltl/sessions-backend/services/data_aggregator/internal/config"
	authhandler "github.com/jonashiltl/sessions-backend/services/data_aggregator/internal/handler/auth_handler"
	partyhandler "github.com/jonashiltl/sessions-backend/services/data_aggregator/internal/handler/party_handler"
	profilehandler "github.com/jonashiltl/sessions-backend/services/data_aggregator/internal/handler/profile_handler"
	storyhandler "github.com/jonashiltl/sessions-backend/services/data_aggregator/internal/handler/story_handler"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	authClient, err := auth.NewClient(c.AUTH_SERVICE_ADDRESS)
	profileClient, err := profile.NewClient(c.PROFILE_SERVICE_ADDRESS)
	partyClient, err := party.NewClient(c.PARTY_SERVICE_ADDRESS)
	storyClient, err := story.NewClient(c.STORY_SERVICE_ADDRESS)

	authHandler := authhandler.NewAuthGatewayHandler(authClient)
	profileHandler := profilehandler.NewProfileGatewayHandler(profileClient)
	partyHandler := partyhandler.NewPartyGatewayHandler(partyClient, profileClient, storyClient)
	storyHandler := storyhandler.NewStoryGatewayHandler(storyClient, profileClient)

	app := fiber.New()

	auth := app.Group("/auth")
	auth.Post("/login", authHandler.Login)
	auth.Post("/register", authHandler.Register)
	auth.Post("/google-login", authHandler.GoogleLogin)

	profile := app.Group("/profile")
	profile.Get("/me", profileHandler.GetMe)
	profile.Get("/:id", profileHandler.GetProfile)
	profile.Get("/:username", profileHandler.GetProfileByUsername)
	profile.Get("/:username", profileHandler.UsernameTaken)
	profile.Patch("/", profileHandler.UpdateProfile)
	// TODO: Add endpoint vor Profile Search

	party := app.Group("/party")
	party.Post("/", partyHandler.CreateParty)
	party.Delete("/:id", partyHandler.DeleteParty)
	party.Get("/:id", partyHandler.GetParty)
	party.Get("/user/:id", partyHandler.GetPartyByUser)
	profile.Patch("/", partyHandler.UpdateParty)

	story := app.Group("/story")
	story.Post("/", storyHandler.CreateStory)
	story.Delete("/:id", storyHandler.DeleteStory)
	story.Get("/:id", storyHandler.GetStory)
	story.Get("/party/:id", storyHandler.GetStoryByParty)
	story.Get("/user/:id", storyHandler.GetStoryByUser)
	story.Get("/presign/:key", storyHandler.PresignURL)

	var sb strings.Builder
	sb.WriteString("0.0.0.0:")
	sb.WriteString(c.PORT)

	fmt.Println("Fiber started at: ", sb.String())
	if err := app.Listen(sb.String()); err != nil {
		log.Fatal(err)
	}
}
