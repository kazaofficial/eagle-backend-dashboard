package repository

import (
	"context"
	"eagle-backend-dashboard/entity"
)

// UserGroupMenuRepository defines the methods for interacting with user group menus.
type UserGroupMenuRepository interface {
	GetUserGroupMenuByParentIDAndUserID(ctx context.Context, parentID int, userID int) ([]entity.UserGroupMenu, error)
	GetByUserGroupIDAndMenuID(ctx context.Context, userGroupID int, menuID int) (*entity.UserGroupMenu, error)
	CreateManyUserGroupMenu(ctx context.Context, userGroupMenus []entity.UserGroupMenu) error
	DeleteManyUserGroupMenuByUserGroupIDAndMenuIDs(ctx context.Context, userGroupID int, menuIDs []int) error
}
