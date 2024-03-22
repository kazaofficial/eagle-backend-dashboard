package controller

import (
	"eagle-backend-dashboard/dto"
	"eagle-backend-dashboard/service"
	"eagle-backend-dashboard/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// UserGroupMenuController is a struct to represent the user group menu controller.
type UserGroupMenuController struct {
	userGroupMenuService service.UserGroupMenuService
}

// NewUserGroupMenuController creates a new instance of UserGroupMenuController.
func NewUserGroupMenuRoutes(handler fiber.Router, userGroupMenuService service.UserGroupMenuService) {
	r := &UserGroupMenuController{
		userGroupMenuService: userGroupMenuService,
	}

	handler.Get("/user-group-menu/:parent_id", r.GetUserGroupMenuByParentIDAndUserID)
	handler.Post("/user-group-menu", r.CreateManyUserGroupMenu)
	handler.Delete("/user-group-menu", r.DeleteManyUserGroupMenu)
}

func (controller *UserGroupMenuController) GetUserGroupMenuByParentIDAndUserID(c *fiber.Ctx) error {
	parentID, err := c.ParamsInt("parent_id")
	id := c.Locals("id").(int)

	if err != nil {
		response := dto.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request, " + err.Error(),
			Error:      nil,
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}
	ctx := c.Context()
	userGroupMenus, err := controller.userGroupMenuService.GetUserGroupMenuByParentIDAndUserID(ctx, parentID, id)
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
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       userGroupMenus,
	}
	return c.Status(http.StatusOK).JSON(response)
}

func (controller *UserGroupMenuController) CreateManyUserGroupMenu(c *fiber.Ctx) error {
	ctx := c.Context()
	request := dto.UserGroupMenuRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		response := dto.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request, " + err.Error(),
			Error:      nil,
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// validate
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

	if len(request.MenuIDs) == 0 {
		response := dto.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request, MenuIDs cannot be empty",
			Error:      nil,
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	err = controller.userGroupMenuService.CreateManyUserGroupMenu(ctx, request)
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
		StatusCode: http.StatusCreated,
		Message:    "Success",
		Data:       nil,
	}
	return c.Status(http.StatusCreated).JSON(response)
}

func (controller *UserGroupMenuController) DeleteManyUserGroupMenu(c *fiber.Ctx) error {
	ctx := c.Context()
	request := dto.UserGroupMenuRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		response := dto.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request, " + err.Error(),
			Error:      nil,
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// validate
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

	if len(request.MenuIDs) == 0 {
		response := dto.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request, MenuIDs cannot be empty",
			Error:      nil,
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	err = controller.userGroupMenuService.DeleteManyUserGroupMenu(ctx, request)
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
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       nil,
	}
	return c.Status(http.StatusOK).JSON(response)
}
