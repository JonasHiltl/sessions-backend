package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jonashiltl/sessions-backend/services/user/ent"
)

type FriendService interface {
	FriendRequest(ctx context.Context, fId uuid.UUID, meId uuid.UUID) error
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
