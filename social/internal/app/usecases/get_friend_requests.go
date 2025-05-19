package usecases

import (
	"context"
	"fmt"

	"github.com/kzemlyak/microchelik/social/internal/app/entities"
)

func (uc *useCase) GetFriendRequests(
	ctx context.Context,
	userId string,
) ([]*entities.FriendRequest, error) {
	var requests []*entities.FriendRequest

	err := uc.TxManager.RunReadCommitted(ctx,
		func(txCtx context.Context) error {
			var err error
			// TODO: Implement GetFriendRequests in repository
			// requests, err = uc.FriendRequestRepository.GetFriendRequests(txCtx, userId)
			if err != nil {
				return fmt.Errorf("%w: %s", ErrGetFriendRequestsFailed, err.Error())
			}

			return nil
		},
	)

	return requests, err
}
