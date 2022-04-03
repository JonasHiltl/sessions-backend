package handler

import (
	ag "github.com/jonashiltl/sessions-backend/packages/grpc/auth"
	"github.com/jonashiltl/sessions-backend/services/auth/internal/service"
)

type authServer struct {
	authService service.AuthService
	token       service.TokenManager
	google      service.GoogleManager
	password    service.PasswordManager
	ag.UnimplementedAuthServiceServer
}

func NewAuthServer(as service.AuthService, tm service.TokenManager, pm service.PasswordManager, google service.GoogleManager) ag.AuthServiceServer {
	return &authServer{
		token:       tm,
		authService: as,
		password:    pm,
		google:      google,
	}
}
