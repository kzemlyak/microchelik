package friend_repository

import (
	"context"

	"github.com/kzemlyak/microchelik/social/internal/app/entities"
)

func (r *FriendRepository) GetFriend(ctx context.Context, id string) (*entities.Friend, error) {
	return nil, nil
}
