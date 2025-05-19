package profile_repository

import (
	"context"

	"github.com/kzemlyak/microchelik/social/internal/app/entities"
)

func (r *ProfileRepository) GetProfile(ctx context.Context, id string) (*entities.Profile, error) {
	return nil, nil
}
