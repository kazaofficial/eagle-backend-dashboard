package entity

import (
	"time"
)

type User struct {
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
