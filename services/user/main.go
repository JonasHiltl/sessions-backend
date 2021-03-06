package main

import (
	"log"
	"sync"

	"github.com/jonashiltl/sessions-backend/packages/stream"
	"github.com/jonashiltl/sessions-backend/services/user/config"
	"github.com/jonashiltl/sessions-backend/services/user/repository"
	"github.com/jonashiltl/sessions-backend/services/user/rpc"
	"github.com/jonashiltl/sessions-backend/services/user/service"
	"github.com/jonashiltl/sessions-backend/services/user/subscribe"
	"github.com/nats-io/nats.go"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	opts := []nats.Option{nats.Name("User Service")}
	nc, err := stream.Connect(c.NATS_CLUSTER, opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer nc.Close()
	stream := stream.New(nc)

	sess, err := repository.NewDB(c.CQL_KEYSPACE, c.CQL_HOSTS)
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	dao := repository.NewDAO(sess)

	upload := service.NewUploadService(c.SPACES_ENDPOINT, c.SPACES_TOKEN)
	token := service.NewTokenManager(c.TOKEN_SECRET)
	google := service.NewGoogleManager(c.GOOGLE_CLIENTID)
	password := service.NewPasswordManager()

	ps := dao.NewProfileRepository()

	s := rpc.NewUserServer(token, google, password, upload, dao.NewUserRepository(), ps, stream)

	var wg sync.WaitGroup
	wg.Add(2)

	go rpc.Start(s, c.PORT)
	go subscribe.NewSubscriber(stream, ps).Start()

	wg.Wait()
}
