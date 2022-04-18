package rpc

import (
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/stream"
	"github.com/jonashiltl/sessions-backend/services/user/internal/service"
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
