package commenthandler

import (
	"github.com/gofiber/fiber/v2"
	cg "github.com/jonashiltl/sessions-backend/packages/grpc/comment"
	uc "github.com/jonashiltl/sessions-backend/packages/grpc/user"
)

type commentGatewayHandler struct {
	cc cg.CommentServiceClient
	uc uc.UserServiceClient
}

type CommentGatewayHandler interface {
	CreateComment(c *fiber.Ctx) error
	DeleteComment(c *fiber.Ctx) error
	GetCommentByParty(c *fiber.Ctx) error
}

func NewCommentGatewayHandler(cc cg.CommentServiceClient, uc uc.UserServiceClient) CommentGatewayHandler {
	return &commentGatewayHandler{
		cc: cc,
		uc: uc,
	}
}
