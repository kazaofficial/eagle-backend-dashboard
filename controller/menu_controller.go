package controller

import (
	"eagle-backend-dashboard/dto"
	"eagle-backend-dashboard/service"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
)

type MenuController struct {
	menuService service.MenuService
}

func NewMenuRoutes(handler fiber.Router, menuService service.MenuService) {
	r := &MenuController{
		menuService: menuService,
	}

	handler.Get("/menu", r.GetMenu)
}

func (controller *MenuController) GetMenu(c *fiber.Ctx) error {
	var request dto.MenuRequest

	ctx := c.Context()
	err := c.QueryParser(&request)
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

	menuResponses, pagination, err := controller.menuService.GetMenu(ctx, &request)
	if err != nil {
		response := dto.ErrorResponse{
			StatusCode: 500,
			Message:    "Internal Server Error",
			Error:      err.Error(),
		}
		return c.Status(500).JSON(response)
	}

	response := dto.ResponseList{
		Response: dto.Response{
			StatusCode: 200,
			Message:    "Success",
			Data:       menuResponses,
		},
		Pagination: pagination,
	}
	return c.Status(200).JSON(response)
}
