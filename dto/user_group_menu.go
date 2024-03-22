package dto

import (
	"time"
)

// UserGroupMenu represents the user_group_menus table in the database.
type UserGroupMenuResponse struct {
	ID          int        `json:"id"`
	UserGroupID int        `json:"user_group_id"`
	MenuID      int        `json:"menu_id"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type UserGroupMenuRequest struct {
	UserGroupID int   `json:"user_group_id" validate:"required"`
	MenuIDs     []int `json:"menu_ids" validate:"required"`
}
