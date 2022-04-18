package rpc

import (
	"log"
	"net"
	"strings"

	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/stream"
	"github.com/jonashiltl/sessions-backend/services/user/internal/service"
	"google.golang.org/grpc"
)

type userServer struct {
	token    service.TokenManager
	google   service.GoogleManager
	password service.PasswordManager
	upload   service.UploadService

	us service.UserService
	ps service.ProfileService

	stream stream.Stream

	ug.UnimplementedUserServiceServer
}

func NewUserServer(
	token service.TokenManager,
	google service.GoogleManager,
	password service.PasswordManager,
	upload service.UploadService,
	us service.UserService,
	ps service.ProfileService,
	stream stream.Stream,
) ug.UserServiceServer {
	return &userServer{
		token:    token,
		google:   google,
		password: password,
		upload:   upload,
		us:       us,
		ps:       ps,
		stream:   stream,
	}
}

func Start(s ug.UserServiceServer, port string) {
	var sb strings.Builder
	sb.WriteString("0.0.0.0:")
	sb.WriteString(port)
	conn, err := net.Listen("tcp", sb.String())
	if err != nil {
		log.Fatalln(err)
	}

	grpc := grpc.NewServer()

	ug.RegisterUserServiceServer(grpc, s)

	log.Println("Starting gRPC Server at: ", sb.String())
	if err := grpc.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
