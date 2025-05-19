package usecases

import (
	"context"
	"fmt"

	"github.com/kzemlyak/microchelik/social/internal/app/entities"
)

func (uc *useCase) GetFriends(
	ctx context.Context,
	userId string,
) ([]*entities.Friend, error) {
	var friends []*entities.Friend

	err := uc.TxManager.RunReadCommitted(ctx,
		func(txCtx context.Context) error {
			var err error
			// TODO: Implement GetFriends in repository
			// friends, err = uc.FriendRepository.GetFriends(txCtx, userId)
			if err != nil {
				return fmt.Errorf("%w: %s", ErrGetFriendsFailed, err.Error())
			}

			return nil
		},
	)

	return friends, err
}
