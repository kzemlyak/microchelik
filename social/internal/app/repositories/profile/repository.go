package profile_repository

import (
	"github.com/kzemlyak/microchelik/social/internal/app/usecases"
	"github.com/kzemlyak/microchelik/social/pkg/postgres/transaction_manager"
)

type ProfileRepository struct {
	txManager transaction_manager.TransactionManagerAPI
}

var _ usecases.ProfileRepository = (*ProfileRepository)(nil)

func NewProfileRepository(txManager transaction_manager.TransactionManagerAPI) *ProfileRepository {
	return &ProfileRepository{
		txManager: txManager,
	}
}
