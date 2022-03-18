package main

import (
	"log"
	"net"

	"github.com/joho/godotenv"
	sg "github.com/jonashiltl/sessions-backend/packages/grpc/story"
	"github.com/jonashiltl/sessions-backend/services/story/internal/handler"
	"github.com/jonashiltl/sessions-backend/services/story/internal/repository"
	"github.com/jonashiltl/sessions-backend/services/story/internal/service"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sess, err := repository.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	dao := repository.NewDAO(&sess)

	sService := service.NewStoryServie(dao)
	us := service.NewUploadService()

	conn, err := net.Listen("tcp", "0.0.0.0:8081")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	sServer := handler.NewStoryServer(sService, us)

	sg.RegisterStoryServiceServer(grpcServer, sServer)

	if err := grpcServer.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
