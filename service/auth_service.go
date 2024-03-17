package service

import (
	"context"
	"eagle-backend-dashboard/dto"
	"eagle-backend-dashboard/repository"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
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
	user, err := s.userRepository.GetUserByUsername(ctx, request.Username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, err
	}

	// create token
	expired_at := time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, dto.Claims{
		ID:          user.ID,
		Username:    user.Username,
		UserGroupID: user.UserGroupID,
		Role:        user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expired_at,
		},
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("APP_SECRET")))
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken: tokenString,
		ExpiredAt:   expired_at,
	}, nil
}
