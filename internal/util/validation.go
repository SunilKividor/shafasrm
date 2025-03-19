package util

import "strings"

const compareEmailString = "@srmist.edu.in"

func ValidateEmail(email string) bool {
	return strings.Contains(email, compareEmailString)
}

func CompareRefreshToken(token_from_user string, token_from_db string) bool {
	return token_from_user == token_from_db
}
