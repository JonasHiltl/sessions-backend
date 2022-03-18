package main

import (
	"fmt"
	"log"
	"net"

	"github.com/joho/godotenv"
	cg "github.com/jonashiltl/sessions-backend/packages/grpc/comment"
	"github.com/jonashiltl/sessions-backend/packages/nats"
	"github.com/jonashiltl/sessions-backend/services/comment/internal/handler"
	"github.com/jonashiltl/sessions-backend/services/comment/internal/repository"
	"github.com/jonashiltl/sessions-backend/services/comment/internal/service"
	gonats "github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	opts := []gonats.Option{gonats.Name("Comment Service")}
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

	cs := service.NewCommentServie(dao, nc)

	addr := "0.0.0.0:8081"

	conn, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	cServer := handler.NewCommentServer(cs)

	cg.RegisterCommentServiceServer(grpcServer, cServer)

	fmt.Println("Starting gRPC Server at: ", addr)
	if err := grpcServer.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
