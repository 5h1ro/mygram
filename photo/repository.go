package photo

import (
	"gorm.io/gorm"
)

type Repository interface {
	Get() ([]Photo, error)
	Find(PhotoID int) (Photo, error)
	Create(photo Photo) (Photo, error)
	Update(photo Photo) (Photo, error)
	Delete(photo Photo) (Photo, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r repository) Get() ([]Photo, error) {
	var photos []Photo
	err := r.db.Find(&photos).Error
	return photos, err
}

func (r repository) Find(PhotoID int) (Photo, error) {
	var photo Photo
	err := r.db.First(&photo, PhotoID).Error
	return photo, err
}

func (r repository) Create(photo Photo) (Photo, error) {
	err := r.db.Create(&photo).Error
	return photo, err
}

func (r repository) Update(photo Photo) (Photo, error) {
	err := r.db.Save(&photo).Error
	return photo, err
}

func (r repository) Delete(photo Photo) (Photo, error) {
	err := r.db.Delete(&photo).Error
	return photo, err
}
