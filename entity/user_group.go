package entity

import "time"

type UserGroup struct {
	ID            int        `json:"id,omitempty"`
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	NumberOfUsers int        `json:"number_of_users"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}

// TableName returns the table name of the UserGroup
func (u *UserGroup) TableName() string {
	return "user_groups"
}
