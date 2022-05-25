package rpc

import (
	"log"
	"net"
	"strings"

	rg "github.com/jonashiltl/sessions-backend/packages/grpc/relation"
	"github.com/jonashiltl/sessions-backend/packages/stream"
	"github.com/jonashiltl/sessions-backend/services/relation/service"
	"google.golang.org/grpc"
)

type relationServer struct {
	fs     service.FriendRelationService
	fp     service.FavoriteParty
	stream stream.Stream
	rg.UnimplementedRelationServiceServer
}

func NewRelationServer(fs service.FriendRelationService, fp service.FavoriteParty, stream stream.Stream) rg.RelationServiceServer {
	return &relationServer{
		fs:     fs,
		fp:     fp,
		stream: stream,
	}
}

func Start(s rg.RelationServiceServer, port string) {
	var sb strings.Builder
	sb.WriteString("0.0.0.0:")
	sb.WriteString(port)
	conn, err := net.Listen("tcp", sb.String())
	if err != nil {
		log.Fatalln(err)
	}

	grpc := grpc.NewServer()

	rg.RegisterRelationServiceServer(grpc, s)

	log.Println("Starting gRPC Server at: ", sb.String())
	if err := grpc.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
