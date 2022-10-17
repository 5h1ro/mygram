package dto

type LoginUser struct {
	Email    string `json:"email" example:"admin@admin.com"`
	Password string `json:"password" example:"password"`
}
