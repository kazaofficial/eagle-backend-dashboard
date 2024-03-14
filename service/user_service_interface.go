package service

import (
	"context"
	"eagle-backend-dashboard/dto"
)

type UserService interface {
	GetUser(ctx context.Context, request *dto.UserListRequest) ([]dto.UserResponse, *dto.Pagination, error)
	GetUserByID(ctx context.Context, id int, me bool) (*dto.UserResponse, error)
	CreateUser(ctx context.Context, request *dto.UserRequest) (*dto.UserResponse, error)
	UpdateUser(ctx context.Context, id int, me bool, request *dto.UserUpdateRequest) (*dto.UserResponse, error)
	DeleteUser(ctx context.Context, id int) (*dto.UserResponse, error)
}
