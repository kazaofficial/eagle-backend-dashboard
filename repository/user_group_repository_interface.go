package repository

import (
	"context"
	"eagle-backend-dashboard/entity"
)

// UserGroupRepository defines the methods for interacting with user groups.
type UserGroupRepository interface {
	GetUserGroup(ctx context.Context, limit *int, offset *int, sort *string, search string) ([]entity.UserGroup, error)
	CountUserGroup(ctx context.Context, search string) (int, error)
	GetUserGroupByID(ctx context.Context, id int) (*entity.UserGroup, error)
	CreateUserGroup(ctx context.Context, userGroup *entity.UserGroup) error
	UpdateUserGroup(ctx context.Context, userGroup *entity.UserGroup) error
	DeleteUserGroup(ctx context.Context, id int) error
}
