package utils

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetStatusCodeFromError(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case gorm.ErrRecordNotFound:
		return http.StatusNotFound
	case bcrypt.ErrMismatchedHashAndPassword:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
