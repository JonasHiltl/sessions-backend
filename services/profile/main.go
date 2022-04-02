package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/joho/godotenv"
	"github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"github.com/jonashiltl/sessions-backend/packages/nats"
	"github.com/jonashiltl/sessions-backend/services/profile/internal/handler"
	"github.com/jonashiltl/sessions-backend/services/profile/internal/repository"
	"github.com/jonashiltl/sessions-backend/services/profile/internal/service"
	gonats "github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	opts := []gonats.Option{gonats.Name("Profile Service")}
	nc, err := nats.Connect(opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer nc.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongo, err := repository.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer mongo.Client().Disconnect(ctx)

	dao := repository.NewDAO(mongo)

	userService := service.NewProfileService(dao)
	uploadService := service.NewUploadService()

	addr := "0.0.0.0:8081"
	conn, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	uServer := handler.NewProfileServer(userService, uploadService)

	profile.RegisterProfileServiceServer(grpcServer, uServer)

	fmt.Println("Starting gRPC Server at: ", addr)
	if err := grpcServer.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
