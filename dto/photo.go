package dto

import "time"

type NewPhotoRequest struct {
	Title     string `json:"title" valid:"required~title cannot be empty" example:"example title"`
	Caption   string `json:"caption" example:"example caption"`
	Photo_Url string `json:"photo_url" valid:"required~Photo URL cannot be empty" example:"http://examplephotourl.jpg"`
	UserID    int    `json:"user_id" example:"1"`
}

type NewPhotoResponse struct {
	Result     string `json:"result" example:"result"`
	Message    string `json:"message" example:"photo have been created"`
	StatusCode int    `json:"statuscode" example:"201"`
}

type PhotoResponse struct {
	ID        int       `json:"id" example:"1"`
	Title     string    `json:"title" example:"example title"`
	Caption   string    `json:"caption" example:"example caption"`
	Photo_Url string    `json:"photo_url" example:"http://examplephotourl.jpg"`
	UserID    int       `json:"user_id" example:"1"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-01"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-01-01"`
}

type GetPhotoResponse struct {
	Result     string          `json:"result" example:"success"`
	Message    string          `json:"message" example:"photo data have been sent"`
	StatusCode int             `json:"statuscode" example:"200"`
	Data       []PhotoResponse `json:"data"`
}
