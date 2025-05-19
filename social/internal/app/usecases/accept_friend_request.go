package usecases

import (
	"context"
	"fmt"

	"github.com/kzemlyak/microchelik/social/internal/app/entities"
)

func (uc *useCase) AcceptFriendRequest(
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

			friendRequest.Status = entities.FriendRequestStatusAccepted
			if err := uc.FriendRequestRepository.UpdateFriendRequest(txCtx, friendRequestId, friendRequest); err != nil {
				return fmt.Errorf("%w: %s", ErrAcceptRequestFailed, err.Error())
			}

			friend1 := entities.NewFriend(friendRequest.SenderID, friendRequest.ReceiverID)
			friend2 := entities.NewFriend(friendRequest.ReceiverID, friendRequest.SenderID)

			if err := uc.FriendRepository.CreateFriend(txCtx, friend1); err != nil {
				return fmt.Errorf("%w: %s", ErrAcceptRequestFailed, err.Error())
			}
			if err := uc.FriendRepository.CreateFriend(txCtx, friend2); err != nil {
				return fmt.Errorf("%w: %s", ErrAcceptRequestFailed, err.Error())
			}

			return nil
		},
	)

	return friendRequest, err
}
