package service

import (
	"context"
	"eagle-backend-dashboard/dto"
	"eagle-backend-dashboard/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
	return &AuthServiceImpl{
		userRepository: userRepository,
	}
}

func (s *AuthServiceImpl) Login(ctx context.Context, request dto.LoginRequest) (*dto.LoginResponse, error) {
	// fmt.Println("-------------", os.Getenv("APP_SECRET"))
	user, err := s.userRepository.GetUserByUsername(ctx, request.Username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, err
	}

	return nil, nil
}
