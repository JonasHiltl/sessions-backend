package main

import (
	"context"
	"log"
	"net"
	"strings"
	"time"

	"github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"github.com/jonashiltl/sessions-backend/packages/stream"
	"github.com/jonashiltl/sessions-backend/services/profile/internal/config"
	"github.com/jonashiltl/sessions-backend/services/profile/internal/repository"
	"github.com/jonashiltl/sessions-backend/services/profile/internal/rpc"
	"github.com/jonashiltl/sessions-backend/services/profile/internal/service"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	opts := []nats.Option{nats.Name("Profile Service")}
	nc, err := stream.Connect(c.NATS_CLUSTER, opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer nc.Close()
	stream := stream.New(nc)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongo, err := repository.NewDB(c.MONGO_URL)
	if err != nil {
		log.Fatal(err)
	}
	defer mongo.Client().Disconnect(ctx)

	dao := repository.NewDAO(mongo)

	userService := service.NewProfileService(dao.NewProfileRepository())
	uploadService := service.NewUploadService(c.SPACES_ENDPOINT, c.SPACES_TOKEN)

	var sb strings.Builder
	sb.WriteString("0.0.0.0:")
	sb.WriteString(c.PORT)
	conn, err := net.Listen("tcp", sb.String())
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	uServer := rpc.NewProfileServer(userService, uploadService, stream)

	profile.RegisterProfileServiceServer(grpcServer, uServer)

	log.Println("Starting gRPC Server at: ", sb.String())
	if err := grpcServer.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
