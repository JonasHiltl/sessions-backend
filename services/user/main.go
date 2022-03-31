package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/joho/godotenv"
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/nats"
	_ "github.com/jonashiltl/sessions-backend/services/user/ent/runtime"
	"github.com/jonashiltl/sessions-backend/services/user/internal/handler"
	"github.com/jonashiltl/sessions-backend/services/user/internal/repository"
	"github.com/jonashiltl/sessions-backend/services/user/internal/service"
	gonats "github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

// @title User Microservice
// @description This Microservice manages user entities
// @version 1.0
// @host localhost:8081
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	opts := []gonats.Option{gonats.Name("User Service")}
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

	_ = service.NewTokenManager()

	userService := service.NewUserService(dao)
	uploadService := service.NewUploadService()
	addr := "0.0.0.0:8081"

	conn, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	uServer := handler.NewUserServer(userService, uploadService)

	ug.RegisterUserServiceServer(grpcServer, uServer)

	fmt.Println("Starting gRPC Server at: ", addr)
	if err := grpcServer.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
