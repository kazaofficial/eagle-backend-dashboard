package entity

import "time"

type UserGroupMenu struct {
	ID          int        `db:"id,omitempty"`
	UserGroupID int        `db:"user_group_id"`
	MenuID      int        `db:"menu_id"`
	CreatedAt   *time.Time `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
}

// TableName returns the table name of the UserGroupMenu
func (u *UserGroupMenu) TableName() string {
	return "user_group_menus"
}
