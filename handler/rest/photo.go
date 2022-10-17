package rest

import (
	"mygram/photo"
	"mygram/photo/dto"
	"mygram/photo/response"
	"mygram/user"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type Photo struct {
	photoService photo.Service
	userService  user.Service
}

func NewPhoto(photoService photo.Service, userService user.Service) *Photo {
	return &Photo{photoService, userService}
}

// GetPhoto godoc
// @Tags photos
// @Description Get photo
// @ID get-photo
// @Accept json
// @Produce json
// @Success 200 {object} []response.PhotoResponse
// @Router /photos [get]
// @Security ApiKeyAuth
func (p *Photo) GetPhoto(c *gin.Context) {

	ud := c.MustGet("userData").(jwt.MapClaims)

	data, e := p.userService.Find(int(ud["id"].(float64)))
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "Token invalid",
		})
		return
	}

	photos, e := p.photoService.Get()
	if e != nil {
		if e.Error() != "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  "photo not found",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  e,
			})
		}
		return
	}

	var photoResponse []response.PhotoResponse

	for _, photo := range photos {

		userResponse := user.Response{
			Email:    data.Email,
			Username: data.Username,
		}

		res := response.PhotoResponse{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoUrl:  photo.PhotoUrl,
			UserID:    photo.UserID,
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
			User:      userResponse,
		}
		photoResponse = append(photoResponse, res)
	}

	c.JSON(http.StatusOK, photoResponse)
}

// CreatePhoto godoc
// @Tags photos
// @Description Create photo
// @ID create-photo
// @Accept json
// @Produce json
// @Param RequestBody body dto.CreatePhoto true "request body json"
// @Success 201 {object} response.PhotoCreateResponse
// @Router /photos [post]
// @Security ApiKeyAuth
func (p *Photo) CreatePhoto(c *gin.Context) {

	user := c.MustGet("userData").(jwt.MapClaims)

	_, e := p.userService.Find(int(user["id"].(float64)))
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "Token invalid",
		})
		return
	}

	var photoRequest dto.CreatePhoto

	rules := govalidator.MapData{
		"title":     []string{"required"},
		"photo_url": []string{"required"},
	}

	opts := govalidator.Options{
		Request: c.Request,
		Data:    &photoRequest,
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

	data, e := p.photoService.Create(int(user["id"].(float64)), photoRequest)

	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  e,
		})
		return
	}
	res := response.PhotoCreateResponse{
		ID:        data.ID,
		Title:     data.Title,
		Caption:   data.Caption,
		PhotoUrl:  data.PhotoUrl,
		UserID:    data.UserID,
		CreatedAt: data.CreatedAt,
	}

	c.JSON(http.StatusCreated, res)
}

// UpdatePhoto godoc
// @Tags photos
// @Description Update photo
// @ID update-photo
// @Accept json
// @Produce json
// @Param photoId path int true "photo id"
// @Param RequestBody body dto.UpdatePhoto true "request body json"
// @Success 200 {object} response.PhotoUpdateResponse
// @Router /photos/{photoId} [put]
// @Security ApiKeyAuth
func (p *Photo) UpdatePhoto(c *gin.Context) {

	user := c.MustGet("userData").(jwt.MapClaims)

	_, e := p.userService.Find(int(user["id"].(float64)))
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "Token invalid",
		})
		return
	}

	var photoRequest dto.UpdatePhoto

	rules := govalidator.MapData{
		"title":     []string{"required"},
		"photo_url": []string{"required"},
		"caption":   []string{"required"},
	}

	opts := govalidator.Options{
		Request: c.Request,
		Data:    &photoRequest,
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

	PhotoID, _ := strconv.Atoi(c.Param("photoId"))
	data, e := p.photoService.Update(int(user["id"].(float64)), PhotoID, photoRequest)

	if e != nil {
		if e.Error() != "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  "photo not found",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  e,
			})
		}
		return
	}
	res := response.PhotoUpdateResponse{
		ID:        data.ID,
		Title:     data.Title,
		Caption:   data.Caption,
		PhotoUrl:  data.PhotoUrl,
		UserID:    data.UserID,
		UpdatedAt: data.UpdatedAt,
	}
	c.JSON(
		http.StatusOK, res)
}

// DeletePhoto godoc
// @Tags photos
// @Description Delete photo
// @ID delete-photo
// @Accept json
// @Produce json
// @Param photoId path int true "photo id"
// @Success 200 {object} response.PhotoDeleteResponse
// @Router /photos/{photoId} [delete]
// @Security ApiKeyAuth
func (p *Photo) DeletePhoto(c *gin.Context) {
	user := c.MustGet("userData").(jwt.MapClaims)

	_, err := p.userService.Find(int(user["id"].(float64)))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "Token invalid",
		})
		return
	}

	PhotoID, _ := strconv.Atoi(c.Param("photoId"))
	_, e := p.photoService.Delete(int(user["id"].(float64)), PhotoID)
	if e != nil {
		if e.Error() != "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  "photo not found",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  e,
			})
		}
		return
	}
	res := response.PhotoDeleteResponse{
		Message: "Your photo has been successfully deleted",
	}
	c.JSON(http.StatusOK, res)
}
