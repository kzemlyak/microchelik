package usecases

import (
	"context"
	"fmt"

	"github.com/kzemlyak/microchelik/social/internal/app/entities"
)

func (uc *useCase) SendFriendRequest(
	ctx context.Context,
	senderId string,
	receiverId string,
) (*entities.FriendRequest, error) {
	friendRequest := entities.NewFriendRequest(senderId, receiverId)

	err := uc.TxManager.RunReadCommitted(ctx,
		func(txCtx context.Context) error {
			// TODO: Add check if friend request already exists
			// TODO: Add check if users are already friends

			if err := uc.FriendRequestRepository.CreateFriendRequest(txCtx, friendRequest); err != nil {
				return fmt.Errorf("%w: %s", ErrSendFriendRequestFailed, err.Error())
			}

			return nil
		},
	)

	return friendRequest, err
}
