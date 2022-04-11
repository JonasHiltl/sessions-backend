package handler

import (
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"github.com/jonashiltl/sessions-backend/packages/stream"
	"github.com/jonashiltl/sessions-backend/services/profile/internal/service"
)

type profileServer struct {
	ps     service.ProfileService
	us     service.UploadService
	stream stream.Stream
	pg.UnimplementedProfileServiceServer
}

func NewProfileServer(ps service.ProfileService, us service.UploadService, stream stream.Stream) pg.ProfileServiceServer {
	return &profileServer{
		us:     us,
		ps:     ps,
		stream: stream,
	}
}
