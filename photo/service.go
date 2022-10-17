package photo

import "mygram/photo/dto"

type Service interface {
	Get() ([]Photo, error)
	Find(PhotoID int) (Photo, error)
	Create(UserID int, req dto.CreatePhoto) (Photo, error)
	Update(UserID int, PhotoID int, req dto.UpdatePhoto) (Photo, error)
	Delete(UserID int, PhotoID int) (Photo, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Get() ([]Photo, error) {
	users, err := s.repository.Get()
	return users, err
}

func (s *service) Find(PhotoID int) (Photo, error) {
	photo, err := s.repository.Find(PhotoID)
	return photo, err
}

func (s *service) Create(UserID int, req dto.CreatePhoto) (Photo, error) {
	photo := Photo{
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoUrl: req.PhotoUrl,
		UserID:   UserID,
	}

	NewPhoto, err := s.repository.Create(photo)
	return NewPhoto, err
}

func (s *service) Update(UserID int, PhotoID int, req dto.UpdatePhoto) (Photo, error) {

	photo, err := s.repository.FindSelf(UserID, PhotoID)

	if err != nil {
		return photo, err
	}

	photo.Title = req.Title
	photo.Caption = req.Caption
	photo.PhotoUrl = req.PhotoUrl

	newPhoto, err := s.repository.Update(photo)
	return newPhoto, err
}

func (s *service) Delete(UserID int, PhotoID int) (Photo, error) {

	photo, err := s.repository.FindSelf(UserID, PhotoID)
	if err != nil {
		return photo, err
	}

	newPhoto, err := s.repository.Delete(photo)
	return newPhoto, err
}
