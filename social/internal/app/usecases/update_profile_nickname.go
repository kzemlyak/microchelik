package usecases

import (
	"context"
	"fmt"

	"github.com/kzemlyak/microchelik/social/internal/app/entities"
)

func (uc *useCase) UpdateProfileNickname(
	ctx context.Context,
	id string,
	nickname string,
) (*entities.Profile, error) {
	var profile *entities.Profile

	err := uc.TxManager.RunReadCommitted(ctx,
		func(txCtx context.Context) error {
			var err error
			profile, err = uc.ProfileRepository.GetProfile(txCtx, id)
			if err != nil {
				return fmt.Errorf("%w: %s", ErrProfileNotFound, err.Error())
			}

			exists := uc.ProfileRepository.DoesProfileExistsByNickname(txCtx, nickname)
			if exists {
				return fmt.Errorf("%w: %s", ErrNicknameAlreadyTaken, nickname)
			}

			profile.Nickname = nickname
			if err := uc.ProfileRepository.UpdateProfile(txCtx, id, profile); err != nil {
				return fmt.Errorf("%w: %s", ErrUpdateProfileFailed, err.Error())
			}

			return nil
		},
	)

	return profile, err
}
