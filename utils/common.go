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

	if strings.Contains(err.Error(), "insert or update") {
		if strings.Contains(err.Error(), "user_group_menus") {
			if strings.Contains(err.Error(), "user_group_menus_menu_id_fkey") {
				return http.StatusNotFound, "Some Menu ID is not found"
			}
			if strings.Contains(err.Error(), "user_group_menus_user_group_id_fkey") {
				return http.StatusNotFound, "User group ID is not found"
			}
		}
	}

	if err.Error() == "All user group menu already exists" {
		return http.StatusConflict, "All user group menu already exists"
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
	case gorm.ErrForeignKeyViolated:
		return http.StatusConflict, "Data is being used with another data"
	default:
		return http.StatusInternalServerError, "Internal Server Error"
	}
}
