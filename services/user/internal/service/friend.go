package service

import (
	"context"

	"github.com/jonashiltl/sessions-backend/services/user/ent"
	"github.com/jonashiltl/sessions-backend/services/user/ent/user"
)

type FriendService interface {
	Get(ctx context.Context, uId string, offset int, limit int) ([]*ent.User, error)
	FriendRequest(ctx context.Context, fId string, meId string) error
	Search(ctx context.Context, uId string, query string, accepted bool) ([]*ent.User, error)
}

type friendSerivce struct {
	client *ent.UserClient
}

func NewFriendService(client *ent.UserClient) FriendService {
	return &friendSerivce{client: client}
}

func (fs *friendSerivce) Get(ctx context.Context, uId string, offset int, limit int) ([]*ent.User, error) {
	friends, err := fs.client.
		Query().
		Where(user.ID(uId)).
		QueryFriends().
		Select(user.FieldID, user.FieldUsername, user.FieldFirstName, user.FieldLastName, user.FieldAvatar, user.FieldRole).
		Offset(offset).
		Limit(limit).
		All(ctx)

	return friends, err
}

func (fs *friendSerivce) FriendRequest(ctx context.Context, fId string, meId string) error {
	_, err := fs.client.
		UpdateOneID(meId).
		AddFriendIDs(fId).
		Save(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (fs *friendSerivce) Search(ctx context.Context, uId string, query string, accepted bool) ([]*ent.User, error) {
	friends, err := fs.client.
		Query().
		Where(user.ID(uId)).
		QueryFriends().
		Where(
			user.Or(
				user.UsernameHasPrefix(query),
				user.FirstNameHasPrefix(query),
				user.LastNameHasPrefix(query),
			),
		).
		Limit(10).
		Select(user.FieldID, user.FieldUsername, user.FieldFirstName, user.FieldLastName, user.FieldAvatar, user.FieldRole).
		All(ctx)

	return friends, err
}
