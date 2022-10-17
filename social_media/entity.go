package social_media

import (
	"time"
)

type SocialMedia struct {
	ID             int    `json:"id" gorm:"type:integer;primaryKey;auto_increment"`
	Name           string `json:"name" gorm:"type:varchar(255)"`
	SocialMediaUrl string `json:"social_media_url" gorm:"type:varchar(255)"`
	UserID         int    `json:"user_id" gorm:"type:integer;not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
