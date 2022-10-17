package comment

import "mygram/comment/dto"

type Service interface {
	Get() ([]Comment, error)
	Create(UserID int, req dto.CreateComment) (Comment, error)
	Update(UserID int, CommentID int, req dto.UpdateComment) (Comment, error)
	Delete(UserID int, CommentID int) (Comment, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Get() ([]Comment, error) {
	comments, err := s.repository.Get()
	return comments, err
}

func (s *service) Create(UserID int, req dto.CreateComment) (Comment, error) {
	comment := Comment{
		Message: req.Message,
		PhotoID: req.PhotoID,
		UserID:  UserID,
	}

	NewPhoto, err := s.repository.Create(comment)
	return NewPhoto, err
}

func (s *service) Update(UserID int, CommentID int, req dto.UpdateComment) (Comment, error) {

	comment, err := s.repository.Find(UserID, CommentID)

	if err != nil {
		return comment, err
	}

	comment.Message = req.Message

	newComment, err := s.repository.Update(comment)
	return newComment, err
}

func (s *service) Delete(UserID int, CommentID int) (Comment, error) {

	comment, err := s.repository.Find(UserID, CommentID)
	if err != nil {
		return comment, err
	}

	newComment, err := s.repository.Delete(comment)
	return newComment, err
}
