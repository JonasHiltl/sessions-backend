package service

import (
	"context"
	"errors"
	"time"

	"github.com/jonashiltl/sessions-backend/services/story/internal/datastruct"
	"github.com/jonashiltl/sessions-backend/services/story/internal/dto"
	"github.com/jonashiltl/sessions-backend/services/story/internal/repository"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/mmcloughlin/geohash"
)

const GEOHASH_PRECISION uint = 9

type StoryService interface {
	Create(c context.Context, s dto.Story) (datastruct.Story, error)
	Delete(c context.Context, uId, sId string) error
	Get(c context.Context, sId string) (datastruct.Story, error)
	GetByUser(c context.Context, uId string) ([]datastruct.Story, error)
	GetByParty(c context.Context, pId string) ([]datastruct.Story, error)
}

type storyService struct {
	dao repository.Dao
}

func NewStoryServie(dao repository.Dao) StoryService {
	return &storyService{dao: dao}
}

func (sService *storyService) Create(c context.Context, s dto.Story) (datastruct.Story, error) {
	nanoid, err := gonanoid.New()
	if err != nil {
		return datastruct.Story{}, errors.New("failed to gen Story id")
	}

	gHash := geohash.EncodeWithPrecision(s.Lat, s.Long, GEOHASH_PRECISION)

	ds := datastruct.Story{
		Id:            nanoid,
		PId:           s.PId,
		UId:           s.UId,
		GHash:         gHash,
		Url:           s.Url,
		TaggedFriends: s.TaggedFriends,
		Created_at:    time.Now(),
	}
	return sService.dao.NewStoryQuery().Create(c, ds)
}

func (sService *storyService) Get(c context.Context, sId string) (datastruct.Story, error) {
	return sService.dao.NewStoryQuery().Get(c, sId)
}

func (sService *storyService) GetByUser(c context.Context, uId string) ([]datastruct.Story, error) {
	return sService.dao.NewStoryQuery().GetByUser(c, uId)
}

func (sService *storyService) GetByParty(c context.Context, pId string) ([]datastruct.Story, error) {
	return sService.dao.NewStoryQuery().GetByParty(c, pId)
}

func (sService *storyService) Delete(c context.Context, uId, sId string) error {
	return sService.dao.NewStoryQuery().Delete(c, uId, sId)
}