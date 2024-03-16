package dto

import "time"

type MenuResponse struct {
	ID          int            `json:"id"`
	ParentID    *int           `json:"parent_id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Icon        string         `json:"icon"`
	Url         string         `json:"url"`
	SubMenus    []MenuResponse `json:"sub_menus"`
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
}

type MenuRequest struct {
	Page  *int   `json:"page"`
	Limit *int   `json:"limit"`
	Sort  string `json:"sort"`
}
