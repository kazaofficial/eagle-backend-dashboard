package repository

import (
	"context"
	"eagle-backend-dashboard/entity"
)

// UserGroupRepository defines the methods for interacting with user groups.
type UserGroupRepository interface {
	GetUserGroup(ctx context.Context, limit *int, offset *int, sort *string) ([]entity.UserGroup, error)
	CountUserGroup(ctx context.Context) (int, error)
	GetUserGroupByID(ctx context.Context, id int) (*entity.UserGroup, error)
}
