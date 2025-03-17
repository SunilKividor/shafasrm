package models

type RegisterBody struct {
	Name       string `json:"name"`
	Username   string `json:"username"`
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
