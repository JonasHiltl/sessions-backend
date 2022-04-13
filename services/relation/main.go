package main

import (
	"log"
	"net"
	"strings"

	rg "github.com/jonashiltl/sessions-backend/packages/grpc/relation"
	"github.com/jonashiltl/sessions-backend/packages/nats"
	"github.com/jonashiltl/sessions-backend/services/relation/internal/config"
	"github.com/jonashiltl/sessions-backend/services/relation/internal/repository"
	"github.com/jonashiltl/sessions-backend/services/relation/internal/rpc"
	gonats "github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	opts := []gonats.Option{gonats.Name("Relation Service")}
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

	var sb strings.Builder
	sb.WriteString("0.0.0.0:")
	sb.WriteString(c.Port)
	conn, err := net.Listen("tcp", sb.String())
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	rServer := rpc.NewRelationServer(dao.NewFriendRelationRepository())

	rg.RegisterRelationServiceServer(grpcServer, rServer)

	log.Println("Starting gRPC Server at: ", sb.String())
	if err := grpcServer.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
