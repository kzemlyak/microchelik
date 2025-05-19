package friend_request_repository

import (
	"context"

	"github.com/kzemlyak/microchelik/social/internal/app/entities"
)

func (r *FriendRequestRepository) CreateFriendRequest(ctx context.Context, friendRequest *entities.FriendRequest) error {
	return nil
}
