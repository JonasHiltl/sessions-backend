package main

import (
	"log"

	"github.com/jonashiltl/sessions-backend/services/story/internal/config"
	"github.com/jonashiltl/sessions-backend/services/story/internal/repository"
	rpc "github.com/jonashiltl/sessions-backend/services/story/internal/rpc"
	"github.com/jonashiltl/sessions-backend/services/story/internal/service"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	sess, err := repository.NewDB(c.SCYLLA_KEYSPACE, c.SCYLLA_HOSTS)
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	dao := repository.NewDAO(&sess)

	s := service.NewStoryServie(dao.NewStoryRepository())
	us := service.NewUploadService(c.SPACES_KEY, c.SPACES_ENDPOINT, c.SPACES_KEY)

	st := rpc.NewStoryServer(s, us)
	rpc.Start(st, c.PORT)
}
