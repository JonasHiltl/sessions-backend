package handler

import (
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/services/user/internal/service"
)

type userServer struct {
	us      service.UserService
	uploadS service.UploadService
	as      service.AuthService
	fs      service.FriendService
	ug.UnimplementedUserServiceServer
}

func NewUserServer(us service.UserService, uploadS service.UploadService, as service.AuthService, fs service.FriendService) ug.UserServiceServer {
	return &userServer{
		as:      as,
		uploadS: uploadS,
		us:      us,
		fs:      fs,
	}
}
