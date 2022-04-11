package main

import (
	"log"
	"net"
	"strings"

	sg "github.com/jonashiltl/sessions-backend/packages/grpc/story"
	"github.com/jonashiltl/sessions-backend/services/story/internal/config"
	"github.com/jonashiltl/sessions-backend/services/story/internal/repository"
	rpc "github.com/jonashiltl/sessions-backend/services/story/internal/rpc"
	"github.com/jonashiltl/sessions-backend/services/story/internal/service"
	"google.golang.org/grpc"
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

	sService := service.NewStoryServie(dao)
	us := service.NewUploadService(c.SPACES_KEY, c.SPACES_ENDPOINT, c.SPACES_KEY)

	var sb strings.Builder
	sb.WriteString("0.0.0.0:")
	sb.WriteString(c.PORT)
	conn, err := net.Listen("tcp", sb.String())
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	sServer := rpc.NewStoryServer(sService, us)

	sg.RegisterStoryServiceServer(grpcServer, sServer)

	log.Println("Starting gRPC Server at: ", sb.String())
	if err := grpcServer.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
