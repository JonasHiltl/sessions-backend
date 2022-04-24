package main

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	cg "github.com/jonashiltl/sessions-backend/packages/grpc/comment"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
	rg "github.com/jonashiltl/sessions-backend/packages/grpc/relation"
	sg "github.com/jonashiltl/sessions-backend/packages/grpc/story"
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/utils/middleware"
	"github.com/jonashiltl/sessions-backend/services/data_aggregator/internal/config"
	authhandler "github.com/jonashiltl/sessions-backend/services/data_aggregator/internal/handler/auth_handler"
	commenthandler "github.com/jonashiltl/sessions-backend/services/data_aggregator/internal/handler/comment_handler"
	partyhandler "github.com/jonashiltl/sessions-backend/services/data_aggregator/internal/handler/party_handler"
	relationhandler "github.com/jonashiltl/sessions-backend/services/data_aggregator/internal/handler/relation_handler"
	storyhandler "github.com/jonashiltl/sessions-backend/services/data_aggregator/internal/handler/story_handler"
	userhandler "github.com/jonashiltl/sessions-backend/services/data_aggregator/internal/handler/user_handler"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	uc, err := ug.NewClient(c.USER_SERVICE_ADDRESS)
	if err != nil {
		log.Fatalf("did not connect to user service: %v", err)
	}
	pc, err := pg.NewClient(c.PARTY_SERVICE_ADDRESS)
	if err != nil {
		log.Fatalf("did not connect to party service: %v", err)
	}
	sc, err := sg.NewClient(c.STORY_SERVICE_ADDRESS)
	if err != nil {
		log.Fatalf("did not connect to story service: %v", err)
	}
	rc, err := rg.NewClient(c.RELATION_SERVICE_ADDRESS)
	if err != nil {
		log.Fatalf("did not connect to relation service: %v", err)
	}
	cc, err := cg.NewClient(c.COMMENT_SERVICE_ADDRESS)
	if err != nil {
		log.Fatalf("did not connect to comment service: %v", err)
	}

	authHandler := authhandler.NewAuthGatewayHandler(uc)
	profileHandler := userhandler.NewUserGatewayHandler(uc, rc)
	partyHandler := partyhandler.NewPartyGatewayHandler(pc, uc, sc)
	storyHandler := storyhandler.NewStoryGatewayHandler(sc, uc)
	relationHandler := relationhandler.NewRelationGatewayHandler(rc)
	commentHandler := commenthandler.NewCommentGatewayHandler(cc, uc)

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
	profile.Get("/me", middleware.AuthRequired(c.TOKEN_SECRET), profileHandler.GetMe)
	profile.Get("/:id", middleware.AuthOptional(c.TOKEN_SECRET), profileHandler.GetProfile)
	profile.Get("/username-taken/:username", profileHandler.UsernameTaken)

	user := app.Group("/user")
	user.Patch("/", middleware.AuthRequired(c.TOKEN_SECRET), profileHandler.UpdateUser)

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

	friend := app.Group("/friend")
	friend.Put("/request/:id", middleware.AuthRequired(c.TOKEN_SECRET), relationHandler.FriendRequest)
	friend.Put("/accept/:id", middleware.AuthRequired(c.TOKEN_SECRET), relationHandler.AcceptFriend)
	friend.Delete("/:id", middleware.AuthRequired(c.TOKEN_SECRET), relationHandler.RemoveFriend)

	comment := app.Group("/comment")
	comment.Post("/party/:id", middleware.AuthRequired(c.TOKEN_SECRET), commentHandler.CreateComment)
	comment.Get("/party/:id", commentHandler.GetCommentByParty)
	comment.Delete("/:id/party/:pId", middleware.AuthRequired(c.TOKEN_SECRET), commentHandler.DeleteComment)
	comment.Post("/:id/reply", middleware.AuthRequired(c.TOKEN_SECRET), commentHandler.CreateReply)
	comment.Get("/:id/reply", commentHandler.GetReplyByComment)
	comment.Delete("/:id/reply/:rId", middleware.AuthRequired(c.TOKEN_SECRET), commentHandler.DeleteReply)

	var sb strings.Builder
	sb.WriteString("0.0.0.0:")
	sb.WriteString(c.PORT)

	if err := app.Listen(sb.String()); err != nil {
		log.Fatal(err)
	}
}
