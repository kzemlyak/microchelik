package friend_request_repository

import (
	"context"

	"github.com/kzemlyak/microchelik/social/internal/app/entities"
)

func (r *FriendRequestRepository) UpdateFriendRequest(ctx context.Context, id string, friendRequest *entities.FriendRequest) error {
	return nil
}
