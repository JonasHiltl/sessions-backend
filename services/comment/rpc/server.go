package rpc

import (
	"log"
	"net"
	"strings"

	cg "github.com/jonashiltl/sessions-backend/packages/grpc/comment"
	"github.com/jonashiltl/sessions-backend/services/comment/service"
	"google.golang.org/grpc"
)

type commentServer struct {
	cs service.CommentService
	rs service.ReplyService
	cg.UnimplementedCommentServiceServer
}

func NewCommentServer(cs service.CommentService, rs service.ReplyService) cg.CommentServiceServer {
	return &commentServer{
		cs: cs,
		rs: rs,
	}
}

func Start(s cg.CommentServiceServer, port string) {
	var sb strings.Builder
	sb.WriteString("0.0.0.0:")
	sb.WriteString(port)
	conn, err := net.Listen("tcp", sb.String())
	if err != nil {
		log.Fatalln(err)
	}

	grpc := grpc.NewServer()

	cg.RegisterCommentServiceServer(grpc, s)

	log.Println("Starting gRPC Server at: ", sb.String())
	if err := grpc.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
