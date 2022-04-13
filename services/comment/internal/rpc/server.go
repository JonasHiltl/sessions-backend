package rpc

import (
	cg "github.com/jonashiltl/sessions-backend/packages/grpc/comment"
	"github.com/jonashiltl/sessions-backend/services/comment/internal/service"
)

type commentServer struct {
	cs service.CommentService
	cg.UnimplementedCommentServiceServer
}

func NewCommentServer(cs service.CommentService) cg.CommentServiceServer {
	return &commentServer{
		cs: cs,
	}
}
