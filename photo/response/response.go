package response

import (
	"mygram/user"
	"time"
)

type PhotoResponse struct {
	ID        int           `json:"id"`
	Title     string        `json:"title"`
	Caption   string        `json:"caption"`
	PhotoUrl  string        `json:"photo_url"`
	UserID    int           `json:"user_id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	User      user.Response `json:"User"`
}

type PhotoCreateResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
type PhotoUpdateResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PhotoDeleteResponse struct {
	Message string `json:"message"`
}
