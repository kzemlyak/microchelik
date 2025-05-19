package entities

import (
	"time"

	"github.com/google/uuid"
)

type FriendRequest struct {
	ID         string
	SenderID   string
	ReceiverID string
	Status     FriendRequestStatus
	CreatedAt  time.Time
}

type FriendRequestStatus string

const (
	FriendRequestStatusPending  FriendRequestStatus = "pending"
	FriendRequestStatusAccepted FriendRequestStatus = "accepted"
	FriendRequestStatusDeclined FriendRequestStatus = "declined"
)

func NewFriendRequest(senderID, receiverID string) *FriendRequest {
	return &FriendRequest{
		ID:         uuid.New().String(),
		SenderID:   senderID,
		ReceiverID: receiverID,
		Status:     FriendRequestStatusPending,
		CreatedAt:  time.Now(),
	}
}
