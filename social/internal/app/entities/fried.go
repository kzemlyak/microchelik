package entities

import (
	"time"

	"github.com/google/uuid"
)

type Friend struct {
	ID        string
	UserID    string
	FriendID  string
	CreatedAt time.Time
}

func NewFriend(userID, friendID string) *Friend {
	return &Friend{
		ID:        uuid.New().String(),
		UserID:    userID,
		FriendID:  friendID,
		CreatedAt: time.Now(),
	}
}
