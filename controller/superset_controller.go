package controller

import (
	"eagle-backend-dashboard/dto"
	"eagle-backend-dashboard/service"
	"eagle-backend-dashboard/utils"
	"os"

	"github.com/gofiber/fiber/v2"
)

type SupersetController struct {
	supersetService service.SupersetService
}

func NewSupersetRoutes(handler fiber.Router, supersetService service.SupersetService) {
	r := &SupersetController{
		supersetService: supersetService,
	}

	handler.Get("/superset/token", r.HandleSupersetToken)
}

func (r *SupersetController) HandleSupersetToken(c *fiber.Ctx) error {
	supersetAuthRequest := dto.SupersetAuthRequest{
		Username: os.Getenv("SUPERSET_USERNAME"),
		Password: os.Getenv("SUPERSET_PASSWORD"),
		Refresh:  true,
		Provider: "db",
	}

	token, err := r.supersetService.AuthTokenRequest(&supersetAuthRequest)
	if err != nil {
		status_code, message := utils.GetStatusCodeFromError(err)
		response := dto.ErrorResponse{
			StatusCode: status_code,
			Message:    message,
			Error:      err.Error(),
		}
		return c.Status(status_code).JSON(response)
	}

	csrfToken, csrfHeader, err := r.supersetService.GetCsrfToken(token.AccessToken)
	if err != nil {
		status_code, message := utils.GetStatusCodeFromError(err)
		response := dto.ErrorResponse{
			StatusCode: status_code,
			Message:    message,
			Error:      err.Error(),
		}
		return c.Status(status_code).JSON(response)
	}

	gTknReq := dto.SupersetGuestTokenRequest{
		User: dto.SupersetUserRequest{
			Username:  "reactapp",
			FirstName: "react",
			LastName:  "app",
		},
		Resources: []dto.SupersetResourceRequest{
			{Type: "dashboard", Id: os.Getenv("SUPERSET_DASHBOARD_ID")},
		},
		Rls: []dto.SupersetRLSRequest{},
	}

	result, err := r.supersetService.GetGuestToken(token.AccessToken, csrfToken, csrfHeader, &gTknReq)
	if err != nil {
		status_code, message := utils.GetStatusCodeFromError(err)
		response := dto.ErrorResponse{
			StatusCode: status_code,
			Message:    message,
			Error:      err.Error(),
		}
		return c.Status(status_code).JSON(response)
	}

	response := dto.Response{
		StatusCode: 200,
		Message:    "Success",
		Data:       result,
	}

	return c.Status(200).JSON(response)
}
