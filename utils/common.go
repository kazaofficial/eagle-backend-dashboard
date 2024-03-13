package utils

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetStatusCodeFromError(err error) (int, string) {
	if err == nil {
		return http.StatusOK, "Success"
	}

	switch err {
	case gorm.ErrRecordNotFound:
		return http.StatusNotFound, "Data Not Found"
	case bcrypt.ErrMismatchedHashAndPassword:
		return http.StatusUnauthorized, "Password is incorrect"
	case gorm.ErrForeignKeyViolated:
		return http.StatusConflict, "Data is being used with another data"
	default:
		return http.StatusInternalServerError, "Internal Server Error"
	}
}
