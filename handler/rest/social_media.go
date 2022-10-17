package rest

import (
	"mygram/social_media"
	"mygram/social_media/dto"
	"mygram/social_media/response"
	"mygram/user"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SocialMedia struct {
	socialMediaService social_media.Service
	userService        user.Service
}

func NewSocialMedia(socialMediaService social_media.Service, userService user.Service) *SocialMedia {
	return &SocialMedia{socialMediaService, userService}
}

// GetSocialMedia godoc
// @Tags socialmedias
// @Description Get social media
// @ID get-social-media
// @Accept json
// @Produce json
// @Success 200 {object} []response.SocialMediaResponse
// @Router /socialmedias [get]
// @Security ApiKeyAuth
func (s *SocialMedia) GetSocialMedia(c *gin.Context) {

	ud := c.MustGet("userData").(jwt.MapClaims)

	data, _ := s.userService.Find(int(ud["id"].(float64)))

	socialMedias, e := s.socialMediaService.Get()
	if e != nil {
		if e.Error() != "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  "social media not found",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  e,
			})
		}
		return
	}

	var socialMediaResponse []response.SocialMediaResponse

	for _, sm := range socialMedias {

		userResponse := response.SocialMediaUserResponse{
			ID:       data.ID,
			Username: data.Username,
		}

		res := response.SocialMediaResponse{
			ID:             sm.ID,
			Name:           sm.Name,
			SocialMediaUrl: sm.SocialMediaUrl,
			UserID:         sm.UserID,
			CreatedAt:      sm.CreatedAt,
			UpdatedAt:      sm.UpdatedAt,
			User:           userResponse,
		}
		socialMediaResponse = append(socialMediaResponse, res)
	}

	c.JSON(http.StatusOK, gin.H{
		"social_medias": socialMediaResponse,
	})
}

// CreateSocialMedia godoc
// @Tags socialmedias
// @Description Create social media
// @ID create-social-media
// @Accept json
// @Produce json
// @Param RequestBody body dto.SocialMedia true "request body json"
// @Success 201 {object} response.SocialMediaCreateResponse
// @Router /socialmedias [post]
// @Security ApiKeyAuth
func (s *SocialMedia) CreateSocialMedia(c *gin.Context) {

	user := c.MustGet("userData").(jwt.MapClaims)

	var socialMediaRequest dto.SocialMedia

	rules := govalidator.MapData{
		"name":             []string{"required"},
		"social_media_url": []string{"required"},
	}

	opts := govalidator.Options{
		Request: c.Request,
		Data:    &socialMediaRequest,
		Rules:   rules,
	}

	v := govalidator.New(opts)
	validate := v.ValidateJSON()
	err := map[string]interface{}{"validationError": validate}
	if len(validate) != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  err,
		})
		return
	}

	data, e := s.socialMediaService.Create(int(user["id"].(float64)), socialMediaRequest)

	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  e,
		})
		return
	}

	res := response.SocialMediaCreateResponse{
		ID:             data.ID,
		Name:           data.Name,
		SocialMediaUrl: data.SocialMediaUrl,
		UserID:         data.UserID,
		CreatedAt:      data.CreatedAt,
	}

	c.JSON(http.StatusCreated, res)
}

// UpdateSocialMedia godoc
// @Tags socialmedias
// @Description Update social media
// @ID update-social-media
// @Accept json
// @Produce json
// @Param socialMediaId path int true "social media id"
// @Param RequestBody body dto.SocialMedia true "request body json"
// @Success 200 {object} response.SocialMediaUpdateResponse
// @Router /socialmedias/{socialMediaId} [put]
// @Security ApiKeyAuth
func (s *SocialMedia) UpdateSocialMedia(c *gin.Context) {

	var socialMediaRequest dto.SocialMedia

	rules := govalidator.MapData{
		"name":             []string{"required"},
		"social_media_url": []string{"required"},
	}

	opts := govalidator.Options{
		Request: c.Request,
		Data:    &socialMediaRequest,
		Rules:   rules,
	}

	v := govalidator.New(opts)
	validate := v.ValidateJSON()
	err := map[string]interface{}{"validationError": validate}
	if len(validate) != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  err,
		})
		return
	}

	SocialMediaID, _ := strconv.Atoi(c.Param("socialMediaId"))
	data, e := s.socialMediaService.Update(SocialMediaID, socialMediaRequest)

	if e != nil {
		if e.Error() != "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  "social media not found",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  e,
			})
		}
		return
	}
	res := response.SocialMediaUpdateResponse{
		ID:             data.ID,
		Name:           data.Name,
		SocialMediaUrl: data.SocialMediaUrl,
		UserID:         data.UserID,
		UpdatedAt:      data.UpdatedAt,
	}

	c.JSON(http.StatusOK, res)
}

// DeleteSocialMedia godoc
// @Tags socialmedias
// @Description Delete social media
// @ID delete-social-media
// @Accept json
// @Produce json
// @Param socialMediaId path int true "social media id"
// @Success 200 {object} response.SocialMediaDeleteResponse
// @Router /socialmedias/{socialMediaId} [delete]
// @Security ApiKeyAuth
func (s *SocialMedia) DeleteSocialMedia(c *gin.Context) {
	SocialMediaID, _ := strconv.Atoi(c.Param("socialMediaId"))
	_, e := s.socialMediaService.Delete(SocialMediaID)
	if e != nil {
		if e.Error() != "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  "social media not found",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  e,
			})
		}
		return
	}
	res := response.SocialMediaDeleteResponse{
		Message: "Your social media has been successfully deleted",
	}

	c.JSON(http.StatusOK, res)
}
