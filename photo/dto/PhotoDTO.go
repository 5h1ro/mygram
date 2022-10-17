package dto

type CreatePhoto struct {
	Title    string `json:"title" example:"First"`
	Caption  string `json:"caption" example:"This is my first photo"`
	PhotoUrl string `json:"photo_url" example:"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSm1HbGcLg0JpQF8teAsREF9tqJH9FmpODS3FTfNZOk&s"`
}

type UpdatePhoto struct {
	Title    string `json:"title" example:"Second"`
	Caption  string `json:"caption" example:"This is my second photo"`
	PhotoUrl string `json:"photo_url" example:"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSm1HbGcLg0JpQF8teAsREF9tqJH9FmpODS3FTfNZOk&s"`
}
