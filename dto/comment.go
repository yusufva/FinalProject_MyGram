package dto

import "time"

type NewCommentRequest struct {
	UserID  int    `json:"user_id"`
	PhotoID int    `json:"photo_id"`
	Message string `json:"message" valid:"required~message cannot be empty"`
}

type NewCommentResponse struct {
	Result     string `json:"result"`
	Message    string `json:"message"`
	StatusCode int    `json:"statuscode"`
}

type CommentResponse struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	PhotoID   int       `json:"photo_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetCommentResponse struct {
	Result     string            `json:"result"`
	Message    string            `json:"message"`
	StatusCode int               `json:"statuscode"`
	Data       []CommentResponse `json:"data"`
}
