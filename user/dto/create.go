package dto

type CreateUser struct {
	Username string `json:"username" example:"Nurhakiki"`
	Email    string `json:"email" example:"admin@admin.com"`
	Password string `json:"password" example:"password"`
	Age      int    `json:"age" example:"10"`
}
