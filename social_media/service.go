package social_media

import "mygram/social_media/dto"

type Service interface {
	Get() ([]SocialMedia, error)
	Create(UserID int, req dto.SocialMedia) (SocialMedia, error)
	Update(UserID int, SocialMediaID int, req dto.SocialMedia) (SocialMedia, error)
	Delete(UserID int, SocialMediaID int) (SocialMedia, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Get() ([]SocialMedia, error) {
	socialMedias, err := s.repository.Get()
	return socialMedias, err
}

func (s *service) Create(UserID int, req dto.SocialMedia) (SocialMedia, error) {
	sm := SocialMedia{
		Name:           req.Name,
		SocialMediaUrl: req.SocialMediaUrl,
		UserID:         UserID,
	}

	newSocialMedia, err := s.repository.Create(sm)
	return newSocialMedia, err
}

func (s *service) Update(UserID int, SocialMediaID int, req dto.SocialMedia) (SocialMedia, error) {

	socialmedia, err := s.repository.Find(UserID, SocialMediaID)

	if err != nil {
		return socialmedia, err
	}

	socialmedia.Name = req.Name
	socialmedia.SocialMediaUrl = req.SocialMediaUrl

	newSocialMedia, err := s.repository.Update(socialmedia)
	return newSocialMedia, err
}

func (s *service) Delete(UserID int, SocialMediaID int) (SocialMedia, error) {

	socialMedia, err := s.repository.Find(UserID, SocialMediaID)
	if err != nil {
		return socialMedia, err
	}

	newSocialMedia, err := s.repository.Delete(socialMedia)
	return newSocialMedia, err
}
