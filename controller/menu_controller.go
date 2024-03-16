package controller

import (
	"eagle-backend-dashboard/dto"
	"eagle-backend-dashboard/service"
	"eagle-backend-dashboard/utils"
	"strconv"

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
	handler.Get("/menu/:id", r.GetMenuByID)
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
		status_code, message := utils.GetStatusCodeFromError(err)
		response := dto.ErrorResponse{
			StatusCode: status_code,
			Message:    message,
			Error:      err.Error(),
		}
		return c.Status(status_code).JSON(response)
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

func (controller *MenuController) GetMenuByID(c *fiber.Ctx) error {
	id := c.Params("id")

	ctx := c.Context()
	idInt, err := strconv.Atoi(id)
	if err != nil {
		response := dto.ErrorResponse{
			StatusCode: 400,
			Message:    "Bad Request, invalid ID",
			Error:      nil,
		}
		return c.Status(400).JSON(response)
	}

	menuResponse, err := controller.menuService.GetMenuByID(ctx, idInt)
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
		Data:       menuResponse,
	}
	return c.Status(200).JSON(response)
}
