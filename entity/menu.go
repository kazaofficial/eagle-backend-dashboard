package entity

import "time"

type Menu struct {
	ID          int        `json:"id" gorm:"primaryKey"`
	ParentID    *int       `json:"parent_id"`
	UrlKey      string     `json:"url_key"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Icon        string     `json:"icon"`
	Url         string     `json:"url"`
	IsShown     bool       `json:"is_shown" gorm:"default:true"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type MenuWithSubMenus struct {
	Menu
	SubMenus []MenuWithSubMenus `json:"sub_menus" gorm:"foreignKey:ParentID"`
}

type MenuWithUserGroup struct {
	Menu
	UserGroupID *int                `json:"user_group_id" gorm:"->"`
	SubMenus    []MenuWithUserGroup `json:"sub_menus" gorm:"foreignKey:ParentID"`
}

// TableName returns the table name of the Menu
func (m *Menu) TableName() string {
	return "menus"
}
