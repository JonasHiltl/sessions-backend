package handler

import (
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"github.com/jonashiltl/sessions-backend/services/profile/internal/service"
)

type profileServer struct {
	us      service.ProfileService
	uploadS service.UploadService
	pg.UnimplementedProfileServiceServer
}

func NewProfileServer(us service.ProfileService, uploadS service.UploadService) pg.ProfileServiceServer {
	return &profileServer{
		uploadS: uploadS,
		us:      us,
	}
}
