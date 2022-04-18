package main

import (
	"log"

	"github.com/jonashiltl/sessions-backend/packages/stream"
	"github.com/jonashiltl/sessions-backend/services/relation/internal/config"
	"github.com/jonashiltl/sessions-backend/services/relation/internal/repository"
	"github.com/jonashiltl/sessions-backend/services/relation/internal/rpc"
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

	sess, err := repository.NewDB(c.SCYLLA_KEYSPACE, c.SCYLLA_HOSTS)
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	dao := repository.NewDAO(&sess)

	r := rpc.NewRelationServer(dao.NewFriendRelationRepository(), stream)
	rpc.Start(r, c.PORT)
}
