package main

import (
	"fmt"
	"log"
	"net"

	"github.com/joho/godotenv"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
	"github.com/jonashiltl/sessions-backend/packages/nats"
	"github.com/jonashiltl/sessions-backend/services/party/internal/handler"
	"github.com/jonashiltl/sessions-backend/services/party/internal/repository"
	"github.com/jonashiltl/sessions-backend/services/party/internal/service"
	gonats "github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	opts := []gonats.Option{gonats.Name("Party Service")}
	nc, err := nats.Connect(opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer nc.Close()

	sess, err := repository.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	dao := repository.NewDAO(&sess)

	partyService := service.NewPartyServie(dao, nc)

	addr := "0.0.0.0:8081"

	conn, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	pServer := handler.NewPartyServer(partyService)

	pg.RegisterPartyServiceServer(grpcServer, pServer)

	fmt.Println("Starting gRPC Server at: ", addr)
	if err := grpcServer.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
