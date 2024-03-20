package utils

import (
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetStatusCodeFromError(err error) (int, string) {
	if err == nil {
		return http.StatusOK, "Success"
	}

	if strings.Contains(err.Error(), "duplicate") {
		if strings.Contains(err.Error(), "users_username_key") {
			return http.StatusConflict, "Username already exists"
		}
	}

	switch err {
	case gorm.ErrRecordNotFound:
		return http.StatusNotFound, "Data Not Found"
	case bcrypt.ErrMismatchedHashAndPassword:
		return http.StatusUnauthorized, "Password is incorrect"
	case gorm.ErrForeignKeyViolated:
		return http.StatusConflict, "Data is being used with another data"
	case gorm.ErrDuplicatedKey:
		return http.StatusConflict, "Data already exists"
	case gorm.ErrRegistered:
		return http.StatusConflict, "Data already exists"
	default:
		return http.StatusInternalServerError, "Internal Server Error"
	}
}
