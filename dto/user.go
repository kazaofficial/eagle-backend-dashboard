package dto

import (
	"time"

	"gorm.io/gorm"
)

type UserResponse struct {
	ID            int                `gorm:"primary_key" json:"id"`
	UserGroupID   int                `json:"user_group_id,omitempty"`
	UserGroup     *UserGroupResponse `json:"user_group,omitempty"`
	Role          string             `json:"role,omitempty"`
	Name          string             `json:"name"`
	Username      string             `json:"username,omitempty"`
	NRP           string             `json:"nrp,omitempty"`
	LastLogin     *time.Time         `json:"last_login,omitempty"`
	CreatedAt     *time.Time         `json:"created_at,omitempty"`
	CreatedBy     int                `json:"created_by,omitempty"`
	CreatedByUser *UserResponse      `json:"created_by_user,omitempty"`
	UpdatedAt     *time.Time         `json:"updated_at,omitempty"`
	UpdatedBy     int                `json:"updated_by,omitempty"`
	DeletedAt     gorm.DeletedAt     `json:"deleted_at,omitempty"`
}

type UserRequest struct {
	UserGroupID int    `json:"user_group_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
	NRP         string `json:"nrp" validate:"required"`
}

type UserUpdateRequest struct {
	UserGroupID int    `json:"user_group_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	NRP         string `json:"nrp" validate:"required"`
}

type UserListRequest struct {
	Page  *int   `json:"page"`
	Limit *int   `json:"limit"`
	Sort  string `json:"sort"`
}
