package models

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID       uuid.UUID   `json:"id"`
	Username string      `json:"username"`
	Follows  []Followers `json:"follows"`
}

type Followers struct {
	FollowerUsername  string `json:"follower_username"`
	FollowingUsername string `json:"following_username"`
}

type Tweets struct {
	ID       uuid.UUID `json:"id"`
	User     string    `json:"user"`
	Content  string    `json:"content"`
	PostedAt time.Time `json:"posted_at"`
}

type CreateResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
