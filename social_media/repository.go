package social_media

import (
	"gorm.io/gorm"
)

type Repository interface {
	Get() ([]SocialMedia, error)
	Find(UserID int, SocialMediaID int) (SocialMedia, error)
	Create(socialMedia SocialMedia) (SocialMedia, error)
	Update(socialMedia SocialMedia) (SocialMedia, error)
	Delete(socialMedia SocialMedia) (SocialMedia, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r repository) Get() ([]SocialMedia, error) {
	var socialMedias []SocialMedia
	err := r.db.Find(&socialMedias).Error
	return socialMedias, err
}

func (r repository) Find(UserID int, SocialMediaID int) (SocialMedia, error) {
	var socialMedia SocialMedia
	err := r.db.Where("user_id = ?", UserID).First(&socialMedia, SocialMediaID).Error
	return socialMedia, err
}

func (r repository) Create(socialMedia SocialMedia) (SocialMedia, error) {
	err := r.db.Create(&socialMedia).Error
	return socialMedia, err
}

func (r repository) Update(socialMedia SocialMedia) (SocialMedia, error) {
	err := r.db.Save(&socialMedia).Error
	return socialMedia, err
}

func (r repository) Delete(socialMedia SocialMedia) (SocialMedia, error) {
	err := r.db.Delete(&socialMedia).Error
	return socialMedia, err
}
