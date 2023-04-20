package dto

import "time"

type NewSocialMediaRequest struct {
	Name             string `json:"name" valid:"required~name cannot be empty" example:"example name"`
	Social_Media_URL string `json:"social_media_url" valid:"required~social media url cannot be empty" example:"https://examplesocmedurl.org"`
	UserID           int    `json:"user_id" example:"1"`
}

type NewSocialMediaResponse struct {
	Result     string `json:"result" example:"success"`
	Message    string `json:"message" example:"social media have been created"`
	StatusCode int    `json:"statusCode" example:"201"`
}

type SocialMediaResponse struct {
	ID               int       `json:"id" example:"1"`
	Name             string    `json:"name" example:"example name"`
	Social_Media_URL string    `json:"social_media_url" example:"https://examplesocmedurl.org"`
	UserID           int       `json:"user_id" example:"1"`
	CreatedAt        time.Time `json:"created_at" example:"2023-01-01"`
	UpdatedAt        time.Time `json:"updated_at" example:"2023-01-01"`
}

type GetSocialMediaResponse struct {
	Result     string                `json:"result" example:"success"`
	Message    string                `json:"message" example:"social media data have been sent"`
	StatusCode int                   `json:"statuscode" example:"200"`
	Data       []SocialMediaResponse `json:"data"`
}
