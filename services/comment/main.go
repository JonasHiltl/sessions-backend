package main

import (
	"log"

	"github.com/jonashiltl/sessions-backend/packages/stream"
	"github.com/jonashiltl/sessions-backend/services/comment/internal/config"
	"github.com/jonashiltl/sessions-backend/services/comment/internal/repository"
	"github.com/jonashiltl/sessions-backend/services/comment/internal/rpc"
	"github.com/jonashiltl/sessions-backend/services/comment/internal/service"
	"github.com/nats-io/nats.go"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	opts := []nats.Option{nats.Name("Comment Service")}
	nc, err := stream.Connect(c.NATS_CLUSTER, opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer nc.Close()

	sess, err := repository.NewDB(c.SCYLLA_KEYSPACE, c.SCYLLA_HOSTS)
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	dao := repository.NewDAO(&sess)

	cs := service.NewCommentService(dao.NewCommentRepository())
	rs := service.NewReplyService(dao.NewReplyRepository())

	s := rpc.NewCommentServer(cs, rs)
	rpc.Start(s, c.PORT)
}
