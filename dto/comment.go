package dto

import "time"

type NewCommentRequest struct {
	UserID  int    `json:"user_id" example:"1"`
	PhotoID int    `json:"photo_id" example:"1"`
	Message string `json:"message" valid:"required~message cannot be empty" example:"example comment"`
}

type NewCommentResponse struct {
	Result     string `json:"result" example:"success"`
	Message    string `json:"message" example:"comment have been created"`
	StatusCode int    `json:"statuscode" example:"201"`
}

type CommentResponse struct {
	ID        int       `json:"id" example:"1"`
	UserID    int       `json:"user_id" example:"1"`
	PhotoID   int       `json:"photo_id" example:"1"`
	Message   string    `json:"message" example:"example comment"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-01"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-01-01"`
}

type GetCommentResponse struct {
	Result     string            `json:"result" example:"success"`
	Message    string            `json:"message" example:"comment data have been sent"`
	StatusCode int               `json:"statuscode" example:"200"`
	Data       []CommentResponse `json:"data"`
}
