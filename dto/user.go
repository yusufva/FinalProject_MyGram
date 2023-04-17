package dto

type NewUserRequest struct {
	Username string `json:"username" valid:"required~username cannot be empty"`
	Email    string `json:"email" valid:"required~email cannot be empty,email"`
	Password string `json:"password" valid:"required~password cannot be empty,stringlength(6|72)"`
	Age      int    `json:"age" valid:"required~age cannot be empty,min=8"`
}

type NewUserResponse struct {
	Result     string `json:"result"`
	StatusCode int    `json:"statuscode"`
	Message    string `json:"message"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type LoginResponse struct {
	Result     string        `json:"result"`
	StatusCode int           `json:"statusCode"`
	Message    string        `json:"message"`
	Data       TokenResponse `json:"data"`
}
