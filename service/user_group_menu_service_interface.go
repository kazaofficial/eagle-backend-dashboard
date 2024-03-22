package service

import (
	"context"
	"eagle-backend-dashboard/dto"
)

// UserGroupMenuService defines the methods for interacting with user group menus.
type UserGroupMenuService interface {
	GetUserGroupMenuByParentIDAndUserID(ctx context.Context, parentID int, userID int) ([]dto.UserGroupMenuResponse, error)
	CreateManyUserGroupMenu(ctx context.Context, request dto.UserGroupMenuRequest) error
	DeleteManyUserGroupMenu(ctx context.Context, request dto.UserGroupMenuRequest) error
}
