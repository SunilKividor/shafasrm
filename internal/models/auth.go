package models

type RegisterRequestBody struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type AuthResBody struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshreqModel struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
