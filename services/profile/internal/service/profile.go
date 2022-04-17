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
	GetMany(ctx context.Context, ids []string) ([]datastruct.Profile, error)
	Update(ctx context.Context, u dto.Profile) (datastruct.Profile, error)
	Delete(ctx context.Context, id string) error
	UsernameTaken(ctx context.Context, uName string) bool
	IncrementFriendCount(ctx context.Context, id string) error
	DecrementFriendCount(ctx context.Context, id string) error
}

type profileService struct {
	repo repository.ProfileRepository
}

func NewProfileService(repo repository.ProfileRepository) ProfileService {
	return &profileService{repo: repo}
}

func (ps *profileService) Create(ctx context.Context, u dto.Profile) (datastruct.Profile, error) {
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

	res, err := ps.repo.Create(ctx, newU)
	if err != nil {
		if strings.Contains(err.Error(), "dup key: { username:") {
			return datastruct.Profile{}, errors.New("username already taken")
		}
	}

	return res, err
}

func (ps *profileService) GetById(ctx context.Context, id string) (datastruct.Profile, error) {
	return ps.repo.GetById(ctx, id)
}
func (ps *profileService) GetByUsername(ctx context.Context, username string) (datastruct.Profile, error) {
	return ps.repo.GetByUsername(ctx, username)
}
func (ps *profileService) GetMany(ctx context.Context, ids []string) ([]datastruct.Profile, error) {
	return ps.repo.GetMany(ctx, ids)
}

func (ps *profileService) Update(ctx context.Context, u dto.Profile) (datastruct.Profile, error) {
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

	return ps.repo.Update(ctx, newU)
}

func (ps *profileService) Delete(ctx context.Context, id string) error {
	return ps.repo.Delete(ctx, id)
}

func (ps *profileService) UsernameTaken(ctx context.Context, uName string) bool {
	return ps.repo.UsernameTaken(ctx, uName)
}

func (ps *profileService) IncrementFriendCount(ctx context.Context, id string) error {
	return ps.repo.IncrementFriendCount(ctx, id)
}

func (ps *profileService) DecrementFriendCount(ctx context.Context, id string) error {
	return ps.repo.DecrementFriendCount(ctx, id)
}
