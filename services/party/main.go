package main

import (
	"log"

	"github.com/jonashiltl/sessions-backend/packages/stream"
	"github.com/jonashiltl/sessions-backend/services/party/internal/config"
	"github.com/jonashiltl/sessions-backend/services/party/internal/repository"
	"github.com/jonashiltl/sessions-backend/services/party/internal/rpc"
	"github.com/jonashiltl/sessions-backend/services/party/internal/service"
	"github.com/nats-io/nats.go"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	opts := []nats.Option{nats.Name("Party Service")}
	nc, err := stream.Connect(c.NATS_CLUSTER, opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer nc.Close()
	stream := stream.New(nc)

	sess, err := repository.NewDB(c.SCYLLA_KEYSPACE, c.SCYLLA_HOSTS)
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	dao := repository.NewDAO(&sess)

	partyService := service.NewPartyServie(dao.NewPartyRepository(), nc)

	p := rpc.NewPartyServer(partyService, stream)
	rpc.Start(p, c.PORT)
}
