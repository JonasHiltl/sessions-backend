package handler

import (
	sg "github.com/jonashiltl/sessions-backend/packages/grpc/story"
	"github.com/jonashiltl/sessions-backend/services/story/internal/service"
)

type storyServer struct {
	sService service.StoryService
	us       service.UploadService
	sg.UnimplementedStoryServiceServer
}

func NewStoryServer(sService service.StoryService, us service.UploadService) sg.StoryServiceServer {
	return &storyServer{sService: sService, us: us}
}
