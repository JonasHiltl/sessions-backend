package main

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/auth"
	"github.com/jonashiltl/sessions-backend/packages/grpc/party"
	"github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"github.com/jonashiltl/sessions-backend/services/data_aggregator/internal/config"
	authhandler "github.com/jonashiltl/sessions-backend/services/data_aggregator/internal/handler/auth_handler"
	partyhandler "github.com/jonashiltl/sessions-backend/services/data_aggregator/internal/handler/party_handler"
	profilehandler "github.com/jonashiltl/sessions-backend/services/data_aggregator/internal/handler/profile_handler"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	ac, err := auth.NewClient(c.AUTH_SERVICE_ADDRESS)
	pc, err := profile.NewClient(c.PROFILE_SERVICE_ADDRESS)
	pac, err := party.NewClient(c.PARTY_SERVICE_ADDRESS)

	ah := authhandler.NewAuthGatewayHandler(ac)
	ph := profilehandler.NewProfileGatewayHandler(pc)
	pah := partyhandler.NewPartyGatewayHandler(pac)

	app := fiber.New()

	auth := app.Group("/auth")
	auth.Post("/login", ah.Login)
	auth.Post("/register", ah.Register)
	auth.Post("/google-login", ah.GoogleLogin)

	var sb strings.Builder
	sb.WriteString("0.0.0.0:")
	sb.WriteString(c.PORT)
	app.Listen(sb.String())
}
