package service

import (
	"context"
	"eagle-backend-dashboard/dto"
)

type UserGroupService interface {
	GetUserGroup(ctx context.Context, request *dto.UserGroupRequest) ([]dto.UserGroupResponse, *dto.Pagination, error)
}
