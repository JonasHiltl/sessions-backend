package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jonashiltl/sessions-backend/services/user/ent"
	"github.com/jonashiltl/sessions-backend/services/user/ent/user"
)

type FriendService interface {
	FriendRequest(ctx context.Context, fId uuid.UUID, meId uuid.UUID) error
	Search(ctx context.Context, userId uuid.UUID, query string, accepted bool) ([]*ent.User, error)
}

type friendSerivce struct {
	client *ent.UserClient
}

func NewFriendService(client *ent.UserClient) FriendService {
	return &friendSerivce{client: client}
}

func (fs *friendSerivce) FriendRequest(ctx context.Context, fId uuid.UUID, meId uuid.UUID) error {
	_, err := fs.client.
		UpdateOneID(meId).
		AddFriendIDs(fId).
		Save(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (fs *friendSerivce) Search(ctx context.Context, userId uuid.UUID, query string, accepted bool) ([]*ent.User, error) {
	friends, err := fs.client.
		Query().
		Where(user.ID(userId)).
		QueryFriends().
		Where(
			user.Or(
				user.UsernameHasPrefix(query),
				user.FirstNameHasPrefix(query),
				user.LastNameHasPrefix(query),
			),
		).
		Limit(5).
		Select(user.FieldID, user.FieldUsername, user.FieldFirstName, user.FieldLastName, user.FieldPicture, user.FieldRole).
		All(ctx)

	return friends, err
}
