package usecases

import (
	"context"
	"errors"
	"fmt"

	"github.com/kzemlyak/microchelik/social/internal/app/entities"
)

var (
	ErrProfileAlreadyExists = errors.New("profile already exists")
	ErrCreateProfileFailed  = errors.New("can't create profile")
)

type CreateProfilePayload struct {
	UserID    string
	Nickname  string
	FirstName *string
	LastName  *string
	Email     *string
}

func (uc *useCase) CreateProfile(
	ctx context.Context,
	payload *CreateProfilePayload,
) (*entities.Profile, error) {
	profile := entities.NewProfile(
		payload.UserID,
		payload.Nickname,
		payload.FirstName,
		payload.LastName,
		payload.Email,
	)

	err := uc.TxManager.RunReadCommitted(ctx,
		func(txCtx context.Context) error {
			exists := uc.ProfileRepository.DoesProfileExistsByNickname(txCtx, payload.Nickname)
			if exists {
				return fmt.Errorf("%w: %s", ErrProfileAlreadyExists, payload.Nickname)
			}

			if err := uc.ProfileRepository.CreateProfile(txCtx, profile); err != nil {
				return fmt.Errorf("%w: %s", ErrCreateProfileFailed, err.Error())
			}

			return nil
		},
	)

	return profile, err
}
