package service

import (
	"context"
	"errors"
	"strings"

	"github.com/jonashiltl/sessions-backend/services/profile/internal/datastruct"
	"github.com/jonashiltl/sessions-backend/services/profile/internal/dto"
	"github.com/jonashiltl/sessions-backend/services/profile/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileService interface {
	Create(ctx context.Context, u dto.Profile) (datastruct.Profile, error)
	GetById(ctx context.Context, id string) (datastruct.Profile, error)
	GetByUsername(ctx context.Context, username string) (datastruct.Profile, error)
	Update(ctx context.Context, u dto.Profile) (datastruct.Profile, error)
	Delete(ctx context.Context, id string) error
	UsernameTaken(ctx context.Context, uName string) bool
}

type profileService struct {
	dao repository.Dao
}

func NewProfileService(dao repository.Dao) ProfileService {
	return &profileService{dao: dao}
}

func (us *profileService) Create(ctx context.Context, u dto.Profile) (datastruct.Profile, error) {
	id := primitive.NewObjectID()

	newU := datastruct.Profile{
		Id:        id,
		Username:  u.Username,
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
		Avatar:    u.Avatar,
	}

	res, err := us.dao.NewProfileRepository().Create(ctx, newU)
	if err != nil {
		if strings.Contains(err.Error(), "dup key: { username:") {
			return datastruct.Profile{}, errors.New("username 2 already taken")
		}
	}

	return res, err
}

func (us *profileService) GetById(ctx context.Context, id string) (datastruct.Profile, error) {
	return us.dao.NewProfileRepository().GetById(ctx, id)
}
func (us *profileService) GetByUsername(ctx context.Context, username string) (datastruct.Profile, error) {
	return us.dao.NewProfileRepository().GetByUsername(ctx, username)
}

func (us *profileService) Update(ctx context.Context, u dto.Profile) (datastruct.Profile, error) {
	id, err := primitive.ObjectIDFromHex(u.Id)
	if err != nil {
		return datastruct.Profile{}, err
	}

	newU := datastruct.Profile{
		Id:        id,
		Username:  u.Username,
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
		Avatar:    u.Avatar,
	}

	return us.dao.NewProfileRepository().Update(ctx, newU)
}

func (us *profileService) Delete(ctx context.Context, id string) error {
	return us.dao.NewProfileRepository().Delete(ctx, id)
}

func (us *profileService) UsernameTaken(ctx context.Context, uName string) bool {
	return us.dao.NewProfileRepository().UsernameTaken(ctx, uName)
}
