package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"github.com/jonashiltl/sessions-backend/packages/nats"
	"github.com/jonashiltl/sessions-backend/services/profile/internal/config"
	"github.com/jonashiltl/sessions-backend/services/profile/internal/handler"
	"github.com/jonashiltl/sessions-backend/services/profile/internal/repository"
	"github.com/jonashiltl/sessions-backend/services/profile/internal/service"
	gonats "github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	opts := []gonats.Option{gonats.Name("Profile Service")}
	nc, err := nats.Connect(c.NatsCluster, opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer nc.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongo, err := repository.NewDB(c.MongoURL)
	if err != nil {
		log.Fatal(err)
	}
	defer mongo.Client().Disconnect(ctx)

	dao := repository.NewDAO(mongo)

	userService := service.NewProfileService(dao)
	uploadService := service.NewUploadService(c.SpacesEndpoint, c.SpacesToken)

	var sb strings.Builder
	sb.WriteString("0.0.0.0:")
	sb.WriteString(c.Port)
	conn, err := net.Listen("tcp", sb.String())
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	uServer := handler.NewProfileServer(userService, uploadService)

	profile.RegisterProfileServiceServer(grpcServer, uServer)

	fmt.Println("Starting gRPC Server at: ", sb.String())
	if err := grpcServer.Serve(conn); err != nil {
		log.Fatal(err)
	}
}