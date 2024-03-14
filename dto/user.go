package dto

import (
	"time"
)

type UserResponse struct {
	ID        int        `gorm:"primary_key" json:"id"`
	GroupID   int        `json:"group_id"`
	Name      string     `json:"name"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	NRP       string     `json:"nrp"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type UserRequest struct {
	GroupID  int    `json:"group_id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	NRP      string `json:"nrp" validate:"required"`
}

type UserUpdateRequest struct {
	GroupID int    `json:"group_id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	NRP     string `json:"nrp" validate:"required"`
}

type UserListRequest struct {
	Page  *int   `json:"page"`
	Limit *int   `json:"limit"`
	Sort  string `json:"sort"`
}
