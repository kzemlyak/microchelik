package profile_repository

import (
	"context"
)

func (r *ProfileRepository) DoesProfileExistsByNickname(ctx context.Context, nickname string) bool {
	return false
}
