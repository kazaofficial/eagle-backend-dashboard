package entity

import (
	"time"
)

type User struct {
	ID          int        `gorm:"primary_key" json:"id,omitempty"`
	UserGroupID int        `json:"user_group_id"`
	Name        string     `json:"name"`
	Username    string     `json:"username"`
	Password    string     `json:"password"`
	NRP         string     `json:"nrp"`
	CreatedAt   *time.Time `json:"created_at"`
	CreatedBy   int        `json:"created_by"`
	UpdatedAt   *time.Time `json:"updated_at"`
	UpdatedBy   int        `json:"updated_by"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

// tableName returns the table name of the User
func (u *User) TableName() string {
	return "users"
}
