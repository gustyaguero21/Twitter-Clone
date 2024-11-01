package models

import "github.com/google/uuid"

type Users struct {
	ID       uuid.UUID   `json:"id"`
	Username string      `json:"username"`
	Follows  []Followers `json:"follows"`
}

type Followers struct {
	FollowerUsername  string `json:"follower_username"`
	FollowingUsername string `json:"following_username"`
}

type UserResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	User    interface{} `json:"user"`
}
