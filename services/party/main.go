package main

import (
	"log"
	"net"
	"strings"

	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
	"github.com/jonashiltl/sessions-backend/packages/nats"
	"github.com/jonashiltl/sessions-backend/services/party/internal/config"
	"github.com/jonashiltl/sessions-backend/services/party/internal/handler"
	"github.com/jonashiltl/sessions-backend/services/party/internal/repository"
	"github.com/jonashiltl/sessions-backend/services/party/internal/service"
	gonats "github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	opts := []gonats.Option{gonats.Name("Party Service")}
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

	partyService := service.NewPartyServie(dao, nc)

	var sb strings.Builder
	sb.WriteString("0.0.0.0:")
	sb.WriteString(c.Port)
	conn, err := net.Listen("tcp", sb.String())
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	pServer := handler.NewPartyServer(partyService)

	pg.RegisterPartyServiceServer(grpcServer, pServer)

	log.Println("Starting gRPC Server at: ", sb.String())
	if err := grpcServer.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
