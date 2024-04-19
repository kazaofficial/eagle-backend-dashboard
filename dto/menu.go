package dto

import "time"

type MenuResponse struct {
	ID          int            `json:"id"`
	ParentID    *int           `json:"parent_id"`
	UrlKey      string         `json:"url_key,omitempty"`
	Name        string         `json:"name"`
	Description string         `json:"description,omitempty"`
	Icon        string         `json:"icon,omitempty"`
	Url         string         `json:"url,omitempty"`
	SubMenus    []MenuResponse `json:"sub_menus"`
	IsShown     bool           `json:"is_shown,omitempty"`
	IsActive    *bool          `json:"is_active,omitempty"`
	CreatedAt   *time.Time     `json:"created_at,omitempty"`
	UpdatedAt   *time.Time     `json:"updated_at,omitempty"`
}

type MenuRequest struct {
	Page  *int   `json:"page"`
	Limit *int   `json:"limit"`
	Sort  string `json:"sort"`
}
