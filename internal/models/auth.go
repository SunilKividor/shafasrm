package models

type RegisterRequestBody struct {
	Name       string `json:"name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Gender     string `json:"gender"`
	Birthday   string `json:"birthday"`
	Location   string `json:"location"`
	Religion   string `json:"religion"`
	Department string `json:"department"`
	Stream     string `json:"stream"`
	Degree     string `json:"degree"`
}

type AuthResBody struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshreqModel struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
