package friend_repository

import (
	"github.com/kzemlyak/microchelik/social/internal/app/usecases"
	"github.com/kzemlyak/microchelik/social/pkg/postgres/transaction_manager"
)

type FriendRepository struct {
	txManager transaction_manager.TransactionManagerAPI
}

var _ usecases.FriendRepository = (*FriendRepository)(nil)

func NewFriendRepository(txManager transaction_manager.TransactionManagerAPI) *FriendRepository {
	return &FriendRepository{
		txManager: txManager,
	}
}
