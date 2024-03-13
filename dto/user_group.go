package dto

import "time"

type UserGroupResponse struct {
	ID            int        `json:"id,omitempty"`
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	NumberOfUsers int        `json:"number_of_users,omitempty"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}

type UserGroupRequest struct {
	Page  *int   `json:"page"`
	Limit *int   `json:"limit"`
	Sort  string `json:"sort"`
}
