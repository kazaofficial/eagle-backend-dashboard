package controller

import (
	"eagle-backend-dashboard/dto"
	"eagle-backend-dashboard/service"
	"eagle-backend-dashboard/utils"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthRoutes(handler fiber.Router, authService service.AuthService) {
	r := &AuthController{
		authService: authService,
	}

	handler.Post("/login", r.Login)
}

func (controller *AuthController) Login(c *fiber.Ctx) error {
	var request dto.LoginRequest

	ctx := c.Context()
	err := c.BodyParser(&request)
	if err != nil {
		response := dto.ErrorResponse{
			StatusCode: 400,
			Message:    "Bad Request, " + err.Error(),
			Error:      nil,
		}
		return c.Status(400).JSON(response)
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var validationErrorMessages []string
			for _, validationError := range validationErrors {
				validationErrorMessages = append(validationErrorMessages, validationError.Error())
			}

			response := dto.ErrorResponse{
				StatusCode: 400,
				Message:    "Validation Error",
				Error:      validationErrorMessages,
			}
			return c.Status(400).JSON(response)
		}
	}

	loginResponse, err := controller.authService.Login(ctx, request)
	if err != nil {
		status_code, message := utils.GetStatusCodeFromError(err)
		response := dto.ErrorResponse{
			StatusCode: status_code,
			Message:    message,
			Error:      err.Error(),
		}
		return c.Status(500).JSON(response)
	}

	response := dto.Response{
		StatusCode: 200,
		Message:    "Success",
		Data:       loginResponse,
	}
	return c.Status(200).JSON(response)
}
