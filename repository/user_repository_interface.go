package repository

import (
	"context"
	"eagle-backend-dashboard/entity"
)

type UserRepository interface {
	GetUser(ctx context.Context, limit *int, offset *int, sort *string) ([]entity.User, error)
	CountUser(ctx context.Context) (int, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	GetUserByID(ctx context.Context, id int, me bool) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, id int) error
}
