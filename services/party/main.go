package main

import (
	"log"

	"github.com/jonashiltl/sessions-backend/packages/stream"
	"github.com/jonashiltl/sessions-backend/services/party/config"
	"github.com/jonashiltl/sessions-backend/services/party/repository"
	"github.com/jonashiltl/sessions-backend/services/party/rpc"
	"github.com/jonashiltl/sessions-backend/services/party/service"
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

	pool, err := repository.NewPGXPool(c.DB_USER, c.DB_PW, c.DB_NAME, c.DB_HOST, c.DB_PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	q := repository.New(pool)

	s := service.NewPartyService(q)

	p := rpc.NewPartyServer(s, stream)
	rpc.Start(p, c.PORT)
}
