package friend_request_repository

import (
	"context"

	"github.com/kzemlyak/microchelik/social/internal/app/entities"
)

func (r *FriendRequestRepository) GetFriendRequest(ctx context.Context, id string) (*entities.FriendRequest, error) {
	return nil, nil
}
