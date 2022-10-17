package comment

import (
	"time"
)

type Comment struct {
	ID        int    `json:"id" gorm:"type:integer;primaryKey;auto_increment"`
	UserID    int    `json:"user_id" gorm:"type:integer;not null"`
	PhotoID   int    `json:"photo_id" gorm:"type:integer;not null"`
	Message   string `json:"message" gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
