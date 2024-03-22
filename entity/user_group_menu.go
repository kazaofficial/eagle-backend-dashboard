package entity

import "time"

type UserGroupMenu struct {
	ID          int        `gorm:"primary_key,omitempty" json:"id,omitempty"`
	UserGroupID int        `json:"user_group_id"`
	MenuID      int        `json:"menu_id"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

// TableName returns the table name of the UserGroupMenu
func (u *UserGroupMenu) TableName() string {
	return "user_group_menus"
}
