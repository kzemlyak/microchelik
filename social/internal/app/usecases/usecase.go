package usecases

import (
	"context"

	"github.com/kzemlyak/microchelik/social/internal/app/entities"
)

type UseCase interface {
	CreateProfile(ctx context.Context, payload *CreateProfilePayload) (*entities.Profile, error)
	UpdateProfileNickname(ctx context.Context, id string, nickname string) (*entities.Profile, error)
	SendFriendRequest(ctx context.Context, senderId string, receiverId string) (*entities.FriendRequest, error)
	AcceptFriendRequest(ctx context.Context, friendRequestId string) (*entities.FriendRequest, error)
	DeclineFriendRequest(ctx context.Context, friendRequestId string) (*entities.FriendRequest, error)
	GetFriendRequests(ctx context.Context, userId string) ([]*entities.FriendRequest, error)
	GetFriends(ctx context.Context, userId string) ([]*entities.Friend, error)
}

type (
	ProfileRepository interface {
		CreateProfile(ctx context.Context, profile *entities.Profile) error
		GetProfile(ctx context.Context, id string) (*entities.Profile, error)
		DoesProfileExistsByNickname(ctx context.Context, nickname string) bool
		UpdateProfile(ctx context.Context, id string, profile *entities.Profile) error
		DeleteProfile(ctx context.Context, id string) error
	}

	FriendRepository interface {
		CreateFriend(ctx context.Context, friend *entities.Friend) error
		GetFriend(ctx context.Context, id string) (*entities.Friend, error)
		UpdateFriend(ctx context.Context, id string, friend *entities.Friend) error
		DeleteFriend(ctx context.Context, id string) error
	}

	FriendRequestRepository interface {
		CreateFriendRequest(ctx context.Context, friendRequest *entities.FriendRequest) error
		GetFriendRequest(ctx context.Context, id string) (*entities.FriendRequest, error)
		UpdateFriendRequest(ctx context.Context, id string, friendRequest *entities.FriendRequest) error
		DeleteFriendRequest(ctx context.Context, id string) error
	}

	OutboxMessagesRepository interface {
		Add(ctx context.Context, id string) error
	}

	TxManager interface {
		RunReadCommitted(ctx context.Context, f func(txCtx context.Context) error) error
	}
)

type Deps struct {
	ProfileRepository       ProfileRepository
	FriendRepository        FriendRepository
	FriendRequestRepository FriendRequestRepository
	TxManager               TxManager
	Outbox                  OutboxMessagesRepository
}

var _ UseCase = (*useCase)(nil)

type useCase struct {
	Deps
}

func NewUseCase(d Deps) *useCase {
	return &useCase{Deps: d}
}
