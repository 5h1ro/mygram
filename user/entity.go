package user

import (
	"mygram/comment"
	"mygram/photo"
	"mygram/social_media"
	"time"
)

type User struct {
	ID           int                      `json:"id" gorm:"type:integer;primaryKey;auto_increment"`
	Username     string                   `json:"username" gorm:"type:varchar(255);unique"`
	Email        string                   `json:"email" gorm:"type:varchar(255);unique"`
	Password     string                   `json:"password" gorm:"type:varchar(255)"`
	Age          int                      `json:"age" gorm:"type:integer"`
	Photos       []photo.Photo            `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Comments     []comment.Comment        `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SocialMedias social_media.SocialMedia `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
