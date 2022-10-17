package response

import "time"

type SocialMediaUserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type SocialMediaResponse struct {
	ID             int                     `json:"id"`
	Name           string                  `json:"name"`
	SocialMediaUrl string                  `json:"social_media_url"`
	UserID         int                     `json:"user_id"`
	CreatedAt      time.Time               `json:"created_at"`
	UpdatedAt      time.Time               `json:"updated_at"`
	User           SocialMediaUserResponse `json:"User"`
}

type SocialMediaCreateResponse struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type SocialMediaUpdateResponse struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         int       `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type SocialMediaDeleteResponse struct {
	Message string `json:"message"`
}
