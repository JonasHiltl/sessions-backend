package main

import (
	"context"
	"log"
	"time"

	"github.com/jonashiltl/sessions-backend/packages/stream"
	"github.com/jonashiltl/sessions-backend/services/user/internal/config"
	"github.com/jonashiltl/sessions-backend/services/user/internal/repository"
	"github.com/jonashiltl/sessions-backend/services/user/internal/rpc"
	"github.com/jonashiltl/sessions-backend/services/user/internal/service"
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongo, err := repository.NewDB(c.MONGO_URL)
	if err != nil {
		log.Fatal(err)
	}
	defer mongo.Client().Disconnect(ctx)

	dao := repository.NewDAO(mongo)

	upload := service.NewUploadService(c.SPACES_ENDPOINT, c.SPACES_TOKEN)
	token := service.NewTokenManager(c.TOKEN_SECRET)
	google := service.NewGoogleManager(c.GOOGLE_CLIENTID)
	password := service.NewPasswordManager()

	s := rpc.NewUserServer(token, google, password, upload, dao.NewUserRepository(), dao.NewProfileRepository(), stream)
	rpc.Start(s, c.PORT)
}
