package repository

import (
	"context"
	"eagle-backend-dashboard/entity"
)

// UserGroupMenuRepository defines the methods for interacting with user group menus.
type UserGroupMenuRepository interface {
	GetUserGroupMenuByParentIDAndUserID(ctx context.Context, parentID int, userID int) ([]entity.UserGroupMenu, error)
}
