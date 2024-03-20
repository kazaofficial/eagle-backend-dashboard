package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID            int            `gorm:"primary_key" json:"id,omitempty"`
	UserGroupID   int            `json:"user_group_id"`
	UserGroup     *UserGroup     `json:"group" gorm:"foreignKey:UserGroupID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Role          string         `json:"role" gorm:"default:'user'"`
	Name          string         `json:"name"`
	Username      string         `json:"username"`
	Password      string         `json:"password"`
	NRP           string         `json:"nrp"`
	LastLogin     *time.Time     `json:"last_login"`
	CreatedAt     *time.Time     `json:"created_at"`
	CreatedBy     int            `json:"created_by"`
	CreatedByUser *User          `json:"created_by_user" gorm:"foreignKey:CreatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	UpdatedAt     *time.Time     `json:"updated_at"`
	UpdatedBy     int            `json:"updated_by"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}

// tableName returns the table name of the User
func (u *User) TableName() string {
	return "users"
}
