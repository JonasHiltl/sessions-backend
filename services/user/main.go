package main

import (
	"context"
	"log"
	"net"
	"strings"
	"time"

	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/stream"
	"github.com/jonashiltl/sessions-backend/services/user/internal/config"
	"github.com/jonashiltl/sessions-backend/services/user/internal/repository"
	"github.com/jonashiltl/sessions-backend/services/user/internal/rpc"
	"github.com/jonashiltl/sessions-backend/services/user/internal/service"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	opts := []nats.Option{nats.Name("User Service")}
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

	upload := service.NewUploadService(c.SPACES_ENDPOINT, c.SPACES_TOKEN)
	token := service.NewTokenManager(c.TOKEN_SECRET)
	google := service.NewGoogleManager(c.GOOGLE_CLIENTID)
	password := service.NewPasswordManager()

	var sb strings.Builder
	sb.WriteString("0.0.0.0:")
	sb.WriteString(c.PORT)
	conn, err := net.Listen("tcp", sb.String())
	if err != nil {
		log.Fatalln(err)
	}

	grpc := grpc.NewServer()

	s := rpc.NewUserServer(token, google, password, upload, dao.NewUserRepository(), dao.NewProfileRepository(), stream)

	ug.RegisterUserServiceServer(grpc, s)

	log.Println("Starting gRPC Server at: ", sb.String())
	if err := grpc.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
