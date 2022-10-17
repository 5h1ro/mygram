package user

import "time"

type Response struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type RegisterResponse struct {
	Age      int    `json:"age"`
	Email    string `json:"email"`
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type UpdateResponse struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Age       int       `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteResponse struct {
	Message string `json:"message"`
}
