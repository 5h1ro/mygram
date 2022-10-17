package rest

import (
	"mygram/comment"
	"mygram/comment/dto"
	"mygram/comment/response"
	"mygram/photo"
	"mygram/user"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type Comment struct {
	commentService comment.Service
	photoService   photo.Service
	userService    user.Service
}

func NewComment(commentService comment.Service, photoService photo.Service, userService user.Service) *Comment {
	return &Comment{commentService, photoService, userService}
}

// GetComment godoc
// @Tags comments
// @Description Get comment
// @ID get-comment
// @Accept json
// @Produce json
// @Success 200 {object} []response.CommentResponse
// @Router /comments [get]
// @Security ApiKeyAuth
func (cm *Comment) GetComment(c *gin.Context) {

	ud := c.MustGet("userData").(jwt.MapClaims)

	_, e := cm.userService.Find(int(ud["id"].(float64)))
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "Token invalid",
		})
		return
	}

	comments, errr := cm.commentService.Get()
	if errr != nil {
		if e.Error() != "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  "photo not found",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  errr,
			})
		}
		return
	}

	var commentResponse []response.CommentResponse

	for _, comment := range comments {

		user, er := cm.userService.Find(comment.UserID)
		if er != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  "User not found",
			})
			return
		}

		photo, err := cm.photoService.Find(comment.PhotoID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  "Photo not found",
			})
			return
		}

		userResponse := response.CommentUserResponse{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
		}

		photoResponse := response.CommentPhotoResponse{
			ID:       photo.ID,
			Title:    photo.Title,
			Caption:  photo.Caption,
			PhotoUrl: photo.PhotoUrl,
			UserID:   photo.UserID,
		}

		res := response.CommentResponse{
			ID:        comment.ID,
			Message:   comment.Message,
			PhotoID:   comment.PhotoID,
			UserID:    comment.UserID,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
			User:      userResponse,
			Photo:     photoResponse,
		}

		commentResponse = append(commentResponse, res)
	}

	c.JSON(http.StatusOK, commentResponse)
}

// CreateComment godoc
// @Tags comments
// @Description Create comment
// @ID create-comment
// @Accept json
// @Produce json
// @Param RequestBody body dto.CreateComment true "request body json"
// @Success 201 {object} response.CommentCreateResponse
// @Router /comments [post]
// @Security ApiKeyAuth
func (cm *Comment) CreateComment(c *gin.Context) {

	user := c.MustGet("userData").(jwt.MapClaims)

	_, e := cm.userService.Find(int(user["id"].(float64)))
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "Token invalid",
		})
		return
	}

	var commentRequest dto.CreateComment

	rules := govalidator.MapData{
		"message":  []string{"required"},
		"photo_id": []string{"required"},
	}

	opts := govalidator.Options{
		Request: c.Request,
		Data:    &commentRequest,
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

	_, er := cm.photoService.Find(commentRequest.PhotoID)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "Photo not found",
		})
		return
	}

	data, e := cm.commentService.Create(int(user["id"].(float64)), commentRequest)

	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  e,
		})
		return
	}

	res := response.CommentCreateResponse{

		ID:        data.ID,
		Message:   data.Message,
		PhotoID:   data.PhotoID,
		UserID:    data.UserID,
		CreatedAt: data.CreatedAt,
	}

	c.JSON(http.StatusCreated, res)
}

// UpdateComment godoc
// @Tags comments
// @Description Update comment
// @ID update-comment
// @Accept json
// @Produce json
// @Param commentId path int true "comment id"
// @Param RequestBody body dto.UpdateComment true "request body json"
// @Success 200 {object} response.CommentUpdateResponse
// @Router /comments/{commentId} [put]
// @Security ApiKeyAuth
func (cm *Comment) UpdateComment(c *gin.Context) {

	user := c.MustGet("userData").(jwt.MapClaims)

	_, e := cm.userService.Find(int(user["id"].(float64)))
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "Token invalid",
		})
		return
	}

	var commentRequest dto.UpdateComment

	rules := govalidator.MapData{
		"message": []string{"required"},
	}

	opts := govalidator.Options{
		Request: c.Request,
		Data:    &commentRequest,
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

	CommentID, _ := strconv.Atoi(c.Param("commentId"))
	data, e := cm.commentService.Update(int(user["id"].(float64)), CommentID, commentRequest)

	if e != nil {
		if e.Error() != "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  "comment not found",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  e,
			})
		}
		return
	}

	res := response.CommentUpdateResponse{
		ID:        data.ID,
		Message:   data.Message,
		PhotoID:   data.PhotoID,
		UserID:    data.UserID,
		UpdatedAt: data.UpdatedAt,
	}

	c.JSON(http.StatusOK, res)
}

// DeleteComment godoc
// @Tags comments
// @Description Delete comment
// @ID delete-comment
// @Accept json
// @Produce json
// @Param commentId path int true "comment id"
// @Success 200 {object} response.CommentDeleteResponse
// @Router /comments/{commentId} [delete]
// @Security ApiKeyAuth
func (cm *Comment) DeleteComment(c *gin.Context) {
	user := c.MustGet("userData").(jwt.MapClaims)

	_, err := cm.userService.Find(int(user["id"].(float64)))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "Token invalid",
		})
		return
	}

	CommentID, _ := strconv.Atoi(c.Param("commentId"))
	_, e := cm.commentService.Delete(int(user["id"].(float64)), CommentID)
	if e != nil {
		if e.Error() != "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  "comment not found",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  e,
			})
		}
		return
	}

	res := response.CommentDeleteResponse{
		Message: "Your comment has been successfully deleted",
	}

	c.JSON(http.StatusOK, res)
}
