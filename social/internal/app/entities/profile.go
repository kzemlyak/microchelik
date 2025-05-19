package entities

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	ID        string
	UserID    string
	Nickname  string
	FirstName *string
	LastName  *string
	Email     *string
	CreatedAt time.Time
}

func NewProfile(
	userID string,
	nickname string,
	firstName *string,
	lastName *string,
	email *string,
) *Profile {
	return &Profile{
		ID:        uuid.New().String(),
		UserID:    userID,
		Nickname:  nickname,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		CreatedAt: time.Now(),
	}
}
