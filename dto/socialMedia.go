package dto

import "time"

type NewSocialMediaRequest struct {
	Name             string `json:"name" valid:"required~name cannot be empty"`
	Social_Media_URL string `json:"social_media_url" valid:"required~social media url cannot be empty"`
	UserID           int    `json:"user_id"`
}

type NewSocialMediaResponse struct {
	Result     string `json:"result"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

type SocialMediaResponse struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Social_Media_URL string    `json:"social_media_url"`
	UserID           int       `json:"user_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type GetSocialMediaResponse struct {
	Result     string                `json:"result"`
	Message    string                `json:"message"`
	StatusCode int                   `json:"statuscode"`
	Data       []SocialMediaResponse `json:"data"`
}
