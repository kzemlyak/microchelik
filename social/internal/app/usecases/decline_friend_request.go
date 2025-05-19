package usecases

import (
	"context"
	"fmt"

	"github.com/kzemlyak/microchelik/social/internal/app/entities"
)

func (uc *useCase) DeclineFriendRequest(
	ctx context.Context,
	friendRequestId string,
) (*entities.FriendRequest, error) {
	var friendRequest *entities.FriendRequest

	err := uc.TxManager.RunReadCommitted(ctx,
		func(txCtx context.Context) error {
			var err error
			friendRequest, err = uc.FriendRequestRepository.GetFriendRequest(txCtx, friendRequestId)
			if err != nil {
				return fmt.Errorf("%w: %s", ErrFriendRequestNotFound, err.Error())
			}

			friendRequest.Status = entities.FriendRequestStatusDeclined
			if err := uc.FriendRequestRepository.UpdateFriendRequest(txCtx, friendRequestId, friendRequest); err != nil {
				return fmt.Errorf("%w: %s", ErrDeclineRequestFailed, err.Error())
			}

			return nil
		},
	)

	return friendRequest, err
}
