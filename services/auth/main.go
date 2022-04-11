package main

import (
	"context"
	"log"
	"net"
	"strings"
	"time"

	ag "github.com/jonashiltl/sessions-backend/packages/grpc/auth"
	"github.com/jonashiltl/sessions-backend/services/auth/internal/config"
	"github.com/jonashiltl/sessions-backend/services/auth/internal/repository"
	rpc "github.com/jonashiltl/sessions-backend/services/auth/internal/rpc"
	"github.com/jonashiltl/sessions-backend/services/auth/internal/service"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongo, err := repository.NewDB(c.MongoURL)
	if err != nil {
		log.Fatal(err)
	}
	defer mongo.Client().Disconnect(ctx)

	dao := repository.NewDAO(mongo)

	tm := service.NewTokenManager(c.TokenSecret)
	gm := service.NewGoogleManager(c.GoogleClientID)
	pm := service.NewPasswordManager()
	as := service.NewAuthService(dao)

	var sb strings.Builder
	sb.WriteString("0.0.0.0:")
	sb.WriteString(c.Port)
	conn, err := net.Listen("tcp", sb.String())
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	uServer := rpc.NewAuthServer(as, tm, pm, gm)

	ag.RegisterAuthServiceServer(grpcServer, uServer)

	log.Println("Starting gRPC Server at: ", sb.String())
	if err := grpcServer.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
