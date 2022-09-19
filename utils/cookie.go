package utils

import (
	"net/http"
)

func createCookie(value string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = "access_token"
	cookie.Value = value

	return cookie
}
