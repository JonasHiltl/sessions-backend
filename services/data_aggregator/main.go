package main

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/jonashiltl/sessions-backend/packages/grpc/auth"
	"github.com/jonashiltl/sessions-backend/packages/grpc/party"
	"github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"github.com/jonashiltl/sessions-backend/packages/grpc/story"
	"github.com/jonashiltl/sessions-backend/packages/utils/middleware"
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
	if err != nil {
		log.Fatalf("did not connect to auth service: %v", err)
	}
	profileClient, err := profile.NewClient(c.PROFILE_SERVICE_ADDRESS)
	if err != nil {
		log.Fatalf("did not connect to profile service: %v", err)
	}
	partyClient, err := party.NewClient(c.PARTY_SERVICE_ADDRESS)
	if err != nil {
		log.Fatalf("did not connect to party service: %v", err)
	}
	storyClient, err := story.NewClient(c.STORY_SERVICE_ADDRESS)
	if err != nil {
		log.Fatalf("did not connect to story service: %v", err)
	}

	authHandler := authhandler.NewAuthGatewayHandler(authClient, profileClient)
	profileHandler := profilehandler.NewProfileGatewayHandler(profileClient)
	partyHandler := partyhandler.NewPartyGatewayHandler(partyClient, profileClient, storyClient)
	storyHandler := storyhandler.NewStoryGatewayHandler(storyClient, profileClient)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's an fiber.*Error
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			// Send custom error in json format
			return ctx.Status(code).JSON(err)
		},
	})
	app.Use(logger.New())
	app.Get("/dashboard", monitor.New())

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
	party.Post("/", middleware.AuthRequired(c.TOKEN_SECRET), partyHandler.CreateParty)
	party.Delete("/:id", middleware.AuthRequired(c.TOKEN_SECRET), partyHandler.DeleteParty)
	party.Get("/:id", partyHandler.GetParty)
	party.Get("/user/:id", partyHandler.GetPartyByUser)
	profile.Patch("/", middleware.AuthRequired(c.TOKEN_SECRET), partyHandler.UpdateParty)

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

	if err := app.Listen(sb.String()); err != nil {
		log.Fatal(err)
	}
}
