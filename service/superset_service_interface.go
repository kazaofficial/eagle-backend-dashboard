package service

import (
	"eagle-backend-dashboard/dto"
)

// MenuService defines the methods for interacting with menus.
type SupersetService interface {
	AuthTokenRequest(request *dto.SupersetAuthRequest) (*dto.SupersetTokenResponse, error)
	GetCsrfToken(accessToken string) (string, string, error)
	GetGuestToken(accessToken string, csrfToken string, cookie string, request *dto.SupersetGuestTokenRequest) (string, error)
}
