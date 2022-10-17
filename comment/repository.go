package comment

import "gorm.io/gorm"

type Repository interface {
	Get() ([]Comment, error)
	Find(UserID int, CommentID int) (Comment, error)
	Create(comment Comment) (Comment, error)
	Update(comment Comment) (Comment, error)
	Delete(comment Comment) (Comment, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r repository) Get() ([]Comment, error) {
	var comment []Comment
	err := r.db.Find(&comment).Error
	return comment, err
}

func (r repository) Find(UserID int, CommentID int) (Comment, error) {
	var comment Comment
	err := r.db.Where("user_id = ?", UserID).First(&comment, CommentID).Error
	return comment, err
}

func (r repository) Create(comment Comment) (Comment, error) {
	err := r.db.Create(&comment).Error
	return comment, err
}

func (r repository) Update(comment Comment) (Comment, error) {
	err := r.db.Save(&comment).Error
	return comment, err
}

func (r repository) Delete(comment Comment) (Comment, error) {
	err := r.db.Delete(&comment).Error
	return comment, err
}
