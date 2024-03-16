package service

import (
	"context"
	"eagle-backend-dashboard/dto"
)

// MenuService defines the methods for interacting with menus.
type MenuService interface {
	GetMenuByUserGroupID(ctx context.Context, userGroupID int) ([]dto.MenuResponse, error)
	GetMenuByIDAndUserGroupID(ctx context.Context, id int, userGroupID int) (*dto.MenuResponse, error)
}
