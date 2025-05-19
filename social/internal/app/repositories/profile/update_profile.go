package profile_repository

import (
	"context"

	"github.com/kzemlyak/microchelik/social/internal/app/entities"
)

func (r *ProfileRepository) UpdateProfile(ctx context.Context, id string, profile *entities.Profile) error {
	return nil
}
