package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserGroup struct {
	ID            int            `json:"id,omitempty" gorm:"primaryKey"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	NumberOfUsers *int           `json:"number_of_users,omitempty" gorm:"->"`
	CreatedAt     *time.Time     `json:"created_at"`
	UpdatedAt     *time.Time     `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}

// TableName returns the table name of the UserGroup
func (u *UserGroup) TableName() string {
	return "user_groups"
}
