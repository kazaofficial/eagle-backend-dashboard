package repository

import (
	"context"
	"eagle-backend-dashboard/entity"
)

type UserRepository interface {
	GetUserByUsername(ctx context.Context, username string) (entity.User, error)
}
