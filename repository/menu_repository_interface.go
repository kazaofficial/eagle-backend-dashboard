package repository

import (
	"context"
	"eagle-backend-dashboard/entity"
)

// MenuRepository defines the methods for interacting with menus.
type MenuRepository interface {
	GetMenuByUserGroupID(ctx context.Context, userGroupID int) ([]entity.Menu, error)
	GetMenuByUrlKeyAndUserGroupID(ctx context.Context, urlKey string, userGroupID int) (*entity.Menu, error)
}
