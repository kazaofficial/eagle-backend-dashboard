package repository

import (
	"context"
	"eagle-backend-dashboard/entity"
)

// MenuRepository defines the methods for interacting with menus.
type MenuRepository interface {
	GetMenu(ctx context.Context, limit *int, offset *int, sort *string) ([]entity.Menu, error)
	CountMenu(ctx context.Context) (int, error)
}
