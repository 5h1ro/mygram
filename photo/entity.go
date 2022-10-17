package photo

import (
	"mygram/comment"
	"time"
)

type Photo struct {
	ID        int               `json:"id" gorm:"type:integer;primaryKey;auto_increment"`
	Title     string            `json:"title" gorm:"type:varchar(255)"`
	Caption   string            `json:"caption" gorm:"type:varchar(255)"`
	PhotoUrl  string            `json:"photo_url" gorm:"type:varchar(255)"`
	UserID    int               `json:"user_id" gorm:"type:integer;not null"`
	Comments  []comment.Comment `gorm:"foreignKey:PhotoID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
