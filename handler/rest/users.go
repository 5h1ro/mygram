package rest

import (
	"mygram/helpers"
	"mygram/user"
	"mygram/user/dto"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type User struct {
	userService user.Service
}

func NewUser(userService user.Service) *User {
	return &User{userService}
}

// Register godoc
// @Tags users
// @Description Register user
// @ID register
// @Accept json
// @Produce json
// @Param RequestBody body dto.CreateUser true "request body json"
// @Success 201 {object} user.RegisterResponse
// @Router /users/register [post]
func (u *User) Register(c *gin.Context) {

	var userRequest dto.CreateUser

	rules := govalidator.MapData{
		"username": []string{"required", "unique:users,username"},
		"email":    []string{"required", "email", "unique:users,email"},
		"password": []string{"required", "min:6"},
		"age":      []string{"required", "min:8"},
	}

	opts := govalidator.Options{
		Request: c.Request,
		Data:    &userRequest,
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
	userRequest.Password = helpers.Hash(userRequest.Password)
	data, e := u.userService.Create(userRequest)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  e,
		})
		return
	}
	res := user.RegisterResponse{
		Age:      data.Age,
		Email:    data.Email,
		ID:       data.ID,
		Username: data.Username,
	}
	c.JSON(
		http.StatusCreated, res)
}

// Login godoc
// @Tags users
// @Description Login user
// @ID login
// @Accept json
// @Produce json
// @Param RequestBody body dto.LoginUser true "request body json"
// @Success 200 {object} user.LoginResponse
// @Router /users/login [post]
func (u *User) Login(c *gin.Context) {

	var userRequest dto.LoginUser

	rules := govalidator.MapData{
		"email":    []string{"required", "email"},
		"password": []string{"required", "min:6"},
	}

	opts := govalidator.Options{
		Request: c.Request,
		Data:    &userRequest,
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

	data, e := u.userService.Login(userRequest.Email)

	if e != nil {
		if e.Error() != "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  "account not found",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  e,
			})
		}
		return
	}

	credentials := helpers.Check([]byte(data.Password), []byte(userRequest.Password))
	if credentials {
		token := helpers.Generate(uint(data.ID), data.Email)
		res := user.LoginResponse{
			Token: token,
		}
		c.JSON(
			http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "Email or Password is wrong",
		})
	}
}

func (u *User) FindUser(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("userID"))
	data, e := u.userService.Find(id)

	if e != nil {
		if e.Error() != "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  "account not found",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  e,
			})
		}
		return
	}

	res := user.Response{
		Email:    data.Email,
		Username: data.Username,
	}

	c.JSON(http.StatusOK, res)
}

// UpdateUser godoc
// @Tags users
// @Description Update user
// @ID update-user
// @Accept json
// @Produce json
// @Param RequestBody body dto.UpdateUser true "request body json"
// @Success 200 {object} user.UpdateResponse
// @Router /users [put]
// @Security ApiKeyAuth
func (u *User) UpdateUser(c *gin.Context) {
	var userRequest dto.UpdateUser

	rules := govalidator.MapData{
		"username": []string{"required", "unique:users,username"},
		"email":    []string{"required", "email", "unique:users,email"},
	}

	opts := govalidator.Options{
		Request: c.Request,
		Data:    &userRequest,
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

	data := c.MustGet("userData").(jwt.MapClaims)

	userUpdate, e := u.userService.Update(int(data["id"].(float64)), userRequest)

	if e != nil {
		if e.Error() != "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  "account not found",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  e,
			})
		}
		return
	}
	res := user.UpdateResponse{
		ID:        userUpdate.ID,
		Email:     userUpdate.Email,
		Username:  userUpdate.Username,
		Age:       userUpdate.Age,
		UpdatedAt: userUpdate.UpdatedAt,
	}

	c.JSON(
		http.StatusOK, res)
}

// DeleteUser godoc
// @Tags users
// @Description Delete user
// @ID delete-user
// @Accept json
// @Produce json
// @Success 200 {object} user.DeleteResponse
// @Router /users [delete]
// @Security ApiKeyAuth
func (u *User) DeleteUser(c *gin.Context) {
	data := c.MustGet("userData").(jwt.MapClaims)

	_, e := u.userService.Delete(int(data["id"].(float64)))

	if e != nil {
		if e.Error() != "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  "account not found",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  e,
			})
		}
		return
	}

	res := user.DeleteResponse{
		Message: "Your account has been successfully deleted",
	}

	c.JSON(
		http.StatusOK, res)
}
