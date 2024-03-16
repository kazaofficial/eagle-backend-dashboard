package entity

import "time"

type Menu struct {
	ID          int        `json:"id" gorm:"primaryKey"`
	ParentID    *int       `json:"parent_id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Icon        string     `json:"icon"`
	Url         string     `json:"url"`
	SubMenus    []Menu     `json:"sub_menus" gorm:"foreignKey:ParentID"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

// TableName returns the table name of the Menu
func (m *Menu) TableName() string {
	return "menus"
}
