package user

import (
	"mygram/user/dto"
)

type Service interface {
	Find(UserID int) (User, error)
	Login(Email string) (User, error)
	Create(req dto.CreateUser) (User, error)
	Update(UserID int, req dto.UpdateUser) (User, error)
	Delete(UserID int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Find(UserID int) (User, error) {
	user, err := s.repository.Find(UserID)
	return user, err
}

func (s *service) Login(Email string) (User, error) {
	user, err := s.repository.Login(Email)
	return user, err
}

func (s *service) Create(req dto.CreateUser) (User, error) {
	user := User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Age:      req.Age,
	}

	NewUser, err := s.repository.Create(user)
	return NewUser, err
}

func (s *service) Update(UserID int, req dto.UpdateUser) (User, error) {

	user, err := s.repository.Find(UserID)

	if err != nil {
		return user, err
	}

	user.Username = req.Username
	user.Email = req.Email

	newUser, err := s.repository.Update(user)
	return newUser, err
}

func (s *service) Delete(UserID int) (User, error) {

	user, err := s.repository.Find(UserID)
	if err != nil {
		return user, err
	}

	newUser, err := s.repository.Delete(user)
	return newUser, err
}
