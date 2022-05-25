package main

import (
	"log"

	"github.com/jonashiltl/sessions-backend/packages/stream"
	"github.com/jonashiltl/sessions-backend/services/relation/config"
	"github.com/jonashiltl/sessions-backend/services/relation/repository"
	"github.com/jonashiltl/sessions-backend/services/relation/rpc"
	"github.com/nats-io/nats.go"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	opts := []nats.Option{nats.Name("Relation Service")}
	nc, err := stream.Connect(c.NATS_CLUSTER, opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer nc.Close()
	stream := stream.New(nc)

	aero, err := repository.NewDB(c.CQL_KEYSPACE, c.CQL_HOSTS)

	if err != nil {
		log.Fatal(err)
	}
	defer aero.Close()

	dao := repository.NewDAO(aero)

	r := rpc.NewRelationServer(dao.NewFriendRelationRepository(), dao.NewFavoritePartyRepository(), stream)
	rpc.Start(r, c.PORT)
}
