package utils

import (
	"net/http"
)

func CreateCookie(key, value string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = key
	cookie.Value = value

	return cookie
}
