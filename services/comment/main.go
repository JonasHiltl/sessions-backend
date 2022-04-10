package main

import (
	"log"
	"net"
	"strings"

	cg "github.com/jonashiltl/sessions-backend/packages/grpc/comment"
	"github.com/jonashiltl/sessions-backend/packages/nats"
	"github.com/jonashiltl/sessions-backend/services/comment/internal/config"
	"github.com/jonashiltl/sessions-backend/services/comment/internal/handler"
	"github.com/jonashiltl/sessions-backend/services/comment/internal/repository"
	"github.com/jonashiltl/sessions-backend/services/comment/internal/service"
	gonats "github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	opts := []gonats.Option{gonats.Name("Comment Service")}
	nc, err := nats.Connect(c.NatsCluster, opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer nc.Close()

	sess, err := repository.NewDB(c.ScyllaKeyspace, c.ScyllaHosts)
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	dao := repository.NewDAO(&sess)

	cs := service.NewCommentServie(dao, nc)

	var sb strings.Builder
	sb.WriteString("0.0.0.0:")
	sb.WriteString(c.Port)
	conn, err := net.Listen("tcp", sb.String())
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	cServer := handler.NewCommentServer(cs)

	cg.RegisterCommentServiceServer(grpcServer, cServer)

	log.Println("Starting gRPC Server at: ", sb.String())
	if err := grpcServer.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
