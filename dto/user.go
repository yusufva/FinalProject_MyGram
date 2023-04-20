package dto

type NewUserRequest struct {
	Username string `json:"username" valid:"required~username cannot be empty,alphanum" example:"exampleusername"`
	Email    string `json:"email" valid:"required~email cannot be empty,email" example:"example@mail.com"`
	Password string `json:"password" valid:"required~password cannot be empty,stringlength(6|72)" example:"examplepassword"`
	Age      int    `json:"age" valid:"required~age cannot be empty,range(8|100)" example:"10"`
}

type NewUserResponse struct {
	Result     string `json:"result" example:"success"`
	StatusCode int    `json:"statuscode" example:"201"`
	Message    string `json:"message" example:"user have been created"`
}

type TokenResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImV4YW1wbGVAbWFpbC5jb20iLCJleHAiOjE2ODE5ODU2MDUsImlkIjowLCJ1c2VybmFtZSI6ImV4YW1wbGV1c2VybmFtZSJ9.hNw3rzukG5U52N0mLRLSW9HKVxR1HUZX8qcgRFlgE70"`
}

type LoginRequest struct {
	Username string `json:"username" valid:"required~username cannot be empty,alphanum" example:"exampleusername"`
	Password string `json:"password" valid:"required~password cannot be empty,stringlength(6|72)" example:"examplepassword"`
}

type LoginResponse struct {
	Result     string        `json:"result" example:"success"`
	StatusCode int           `json:"statusCode" example:"200"`
	Message    string        `json:"message" example:"logged in successfully"`
	Data       TokenResponse `json:"data"`
}
