package models

import "github.com/google/uuid"

type Users struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Follows  []int64   `json:"follows"`
}

type Follows struct {
	ID         uuid.UUID `json:"id"`
	FollowerID int64     `json:"follower_id"`
	FollowedID int64     `json:"followed_id"`
}

type Tweets struct {
	ID       uuid.UUID `json:"id"`
	UserID   int64     `json:"user_id"`
	Content  string    `json:"content"`
	PostedAt int64     `json:"posted_at"`
}

type UserResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	User    interface{} `json:"user"`
}
