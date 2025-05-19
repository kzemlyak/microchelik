package friend_request_repository

import (
	"github.com/kzemlyak/microchelik/social/internal/app/usecases"
	"github.com/kzemlyak/microchelik/social/pkg/postgres/transaction_manager"
)

type FriendRequestRepository struct {
	txManager transaction_manager.TransactionManagerAPI
}

var _ usecases.FriendRequestRepository = (*FriendRequestRepository)(nil)

func NewFriendRequestRepository(txManager transaction_manager.TransactionManagerAPI) *FriendRequestRepository {
	return &FriendRequestRepository{
		txManager: txManager,
	}
}
