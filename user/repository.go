package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	Find(UserID int) (User, error)
	Login(Email string) (User, error)
	Create(user User) (User, error)
	Update(user User) (User, error)
	Delete(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r repository) Find(UserID int) (User, error) {
	var user User
	err := r.db.First(&user, UserID).Error
	return user, err
}

func (r repository) Login(Email string) (User, error) {
	var user User
	err := r.db.First(&user, "email = ?", Email).Error
	return user, err
}

func (r repository) Create(user User) (User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func (r repository) Delete(user User) (User, error) {
	err := r.db.Delete(&user).Error
	return user, err
}
