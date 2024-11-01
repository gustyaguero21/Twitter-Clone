package models

import "github.com/google/uuid"

type Users struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Follows  []int64   `json:"follows"`
}

type Follows struct {
	ID         uuid.UUID
	FollowerID int64
	FollowedID int64
}

type Tweets struct {
	ID       uuid.UUID
	UserID   int64
	Content  string
	PostedAt int64
}
