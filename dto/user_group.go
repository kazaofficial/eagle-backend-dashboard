package dto

import (
	"time"

	"gorm.io/gorm"
)

type UserGroupResponse struct {
	ID            int            `json:"id,omitempty"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	NumberOfUsers *int           `json:"number_of_users,omitempty"`
	Menus         []MenuResponse `json:"menus,omitempty"`
	CreatedAt     *time.Time     `json:"created_at"`
	UpdatedAt     *time.Time     `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}

type UserGroupRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type UserGroupListRequest struct {
	Page   *int   `json:"page"`
	Limit  *int   `json:"limit"`
	Sort   string `json:"sort"`
	Search string `json:"search"`
}
