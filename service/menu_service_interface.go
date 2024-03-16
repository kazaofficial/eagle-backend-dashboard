package service

import (
	"context"
	"eagle-backend-dashboard/dto"
)

// MenuService defines the methods for interacting with menus.
type MenuService interface {
	GetMenu(ctx context.Context, request *dto.MenuRequest) ([]dto.MenuResponse, *dto.Pagination, error)
	GetMenuByID(ctx context.Context, id int) (*dto.MenuResponse, error)
}
