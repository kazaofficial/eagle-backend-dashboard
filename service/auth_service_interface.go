package service

import (
	"context"
	"eagle-backend-dashboard/dto"
)

type AuthService interface {
	Login(ctx context.Context, request dto.LoginRequest) (*dto.LoginResponse, error)
}
