package usecases

import "errors"

var (
	// Profile errors
	ErrProfileNotFound      = errors.New("profile not found")
	ErrNicknameAlreadyTaken = errors.New("nickname already taken")
	ErrUpdateProfileFailed  = errors.New("can't update profile")

	// Friend request errors
	ErrFriendRequestNotFound      = errors.New("friend request not found")
	ErrFriendRequestAlreadyExists = errors.New("friend request already exists")
	ErrSendFriendRequestFailed    = errors.New("can't send friend request")
	ErrAcceptRequestFailed        = errors.New("can't accept friend request")
	ErrDeclineRequestFailed       = errors.New("can't decline friend request")
	ErrGetFriendRequestsFailed    = errors.New("can't get friend requests")

	// Friend errors
	ErrGetFriendsFailed = errors.New("can't get friends list")
)
