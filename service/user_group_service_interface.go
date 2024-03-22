package service

import (
	"context"
	"eagle-backend-dashboard/dto"
)

type UserGroupService interface {
	GetUserGroup(ctx context.Context, request *dto.UserGroupListRequest) ([]dto.UserGroupResponse, *dto.Pagination, error)
	GetUserGroupWithAccess(ctx context.Context, request *dto.UserGroupListRequest) ([]dto.UserGroupResponse, *dto.Pagination, error)
	GetUserGroupByID(ctx context.Context, id int) (*dto.UserGroupResponse, error)
	CreateUserGroup(ctx context.Context, request *dto.UserGroupRequest) (*dto.UserGroupResponse, error)
	UpdateUserGroup(ctx context.Context, id int, request *dto.UserGroupRequest) (*dto.UserGroupResponse, error)
	DeleteUserGroup(ctx context.Context, id int) (*dto.UserGroupResponse, error)
}
