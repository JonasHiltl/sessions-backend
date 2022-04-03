package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/joho/godotenv"
	ag "github.com/jonashiltl/sessions-backend/packages/grpc/auth"
	"github.com/jonashiltl/sessions-backend/services/auth/internal/handler"
	"github.com/jonashiltl/sessions-backend/services/auth/internal/repository"
	"github.com/jonashiltl/sessions-backend/services/auth/internal/service"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongo, err := repository.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer mongo.Client().Disconnect(ctx)

	dao := repository.NewDAO(mongo)

	tm := service.NewTokenManager()
	gm := service.NewGoogleManager()
	pm := service.NewPasswordManager()
	as := service.NewAuthService(dao)

	addr := "0.0.0.0:8081"
	conn, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	uServer := handler.NewAuthServer(as, tm, pm, gm)

	ag.RegisterAuthServiceServer(grpcServer, uServer)

	fmt.Println("Starting gRPC Server at: ", addr)
	if err := grpcServer.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
