package dto

type CreateComment struct {
	Message string `json:"message" example:"Nice pict"`
	PhotoID int    `json:"photo_id" example:"1"`
}

type UpdateComment struct {
	Message string `json:"message" example:"Good picture"`
}
