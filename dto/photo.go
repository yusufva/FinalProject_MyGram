package dto

import "time"

type NewPhotoRequest struct {
	Title     string `json:"title" valid:"required~title cannot be empty"`
	Caption   string `json:"caption"`
	Photo_Url string `json:"photo_url" valid:"required~Photo URL cannot be empty"`
	UserID    int    `json:"user_id" valid:"required~User id cannot be empty"`
}

type NewPhotoResponse struct {
	Result     string `json:"result"`
	Message    string `json:"message"`
	StatusCode int    `json:"statuscode"`
}

type PhotoResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	Photo_Url string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetPhotoResponse struct {
	Result     string          `json:"result"`
	Message    string          `json:"message"`
	StatusCode int             `json:"statuscode"`
	Data       []PhotoResponse `json:"data"`
}
