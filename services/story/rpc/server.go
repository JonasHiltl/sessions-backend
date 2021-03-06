package rpc

import (
	"log"
	"net"
	"strings"

	sg "github.com/jonashiltl/sessions-backend/packages/grpc/story"
	"github.com/jonashiltl/sessions-backend/services/story/service"
	"google.golang.org/grpc"
)

type storyServer struct {
	sService service.StoryService
	us       service.UploadService
	sg.UnimplementedStoryServiceServer
}

func NewStoryServer(sService service.StoryService, us service.UploadService) sg.StoryServiceServer {
	return &storyServer{sService: sService, us: us}
}

func Start(s sg.StoryServiceServer, port string) {
	var sb strings.Builder
	sb.WriteString("0.0.0.0:")
	sb.WriteString(port)
	conn, err := net.Listen("tcp", sb.String())
	if err != nil {
		log.Fatalln(err)
	}

	grpc := grpc.NewServer()

	sg.RegisterStoryServiceServer(grpc, s)

	log.Println("Starting gRPC Server at: ", sb.String())
	if err := grpc.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
