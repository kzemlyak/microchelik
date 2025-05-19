package friend_repository

import (
	"context"

	"github.com/kzemlyak/microchelik/social/internal/app/entities"
)

func (r *FriendRepository) UpdateFriend(ctx context.Context, id string, friend *entities.Friend) error {
	return nil
}
