package handler

import (
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/services/user/internal/service"
)

type userServer struct {
	us      service.UserService
	uploadS service.UploadService
	ug.UnimplementedUserServiceServer
}

func NewUserServer(us service.UserService, uploadS service.UploadService) ug.UserServiceServer {
	return &userServer{
		uploadS: uploadS,
		us:      us,
	}
}
