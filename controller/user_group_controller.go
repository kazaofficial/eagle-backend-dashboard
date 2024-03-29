package controller

import (
	"eagle-backend-dashboard/dto"
	"eagle-backend-dashboard/service"
	"eagle-backend-dashboard/utils"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
)

type UserGroupController struct {
	userGroupService service.UserGroupService
}

func NewUserGroupRoutes(handler fiber.Router, userGroupService service.UserGroupService) {
	r := &UserGroupController{
		userGroupService: userGroupService,
	}

	handler.Get("/user-group", r.GetUserGroup)
	handler.Get("/user-group/access", r.GetUserGroupWithAccess)
	handler.Get("/user-group/:id", r.GetUserGroupByID)
	handler.Post("/user-group", r.CreateUserGroup)
	handler.Put("/user-group/:id", r.UpdateUserGroup)
	handler.Delete("/user-group/:id", r.DeleteUserGroup)
}

func (controller *UserGroupController) GetUserGroup(c *fiber.Ctx) error {
	var request dto.UserGroupListRequest

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

	userGroupResponses, pagination, err := controller.userGroupService.GetUserGroup(ctx, &request)
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
			Data:       userGroupResponses,
		},
		Pagination: pagination,
	}

	return c.Status(200).JSON(response)
}

func (controller *UserGroupController) GetUserGroupWithAccess(c *fiber.Ctx) error {
	var request dto.UserGroupListRequest

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

	userGroupResponses, pagination, err := controller.userGroupService.GetUserGroupWithAccess(ctx, &request)
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
			Data:       userGroupResponses,
		},
		Pagination: pagination,
	}

	return c.Status(200).JSON(response)
}

func (controller *UserGroupController) GetUserGroupByID(c *fiber.Ctx) error {
	id := c.Params("id")

	ctx := c.Context()
	idInt, err := strconv.Atoi(id) // Change this line
	if err != nil {
		response := dto.ErrorResponse{
			StatusCode: 400,
			Message:    "Bad Request, " + err.Error(),
			Error:      nil,
		}
		return c.Status(400).JSON(response)
	}
	userGroupResponse, err := controller.userGroupService.GetUserGroupByID(ctx, idInt)
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
		Data:       userGroupResponse,
	}

	return c.Status(200).JSON(response)
}

func (controller *UserGroupController) CreateUserGroup(c *fiber.Ctx) error {
	var request dto.UserGroupRequest

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

	userGroupResponse, err := controller.userGroupService.CreateUserGroup(ctx, &request)
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
		Data:       userGroupResponse,
	}

	return c.Status(200).JSON(response)
}

func (controller *UserGroupController) UpdateUserGroup(c *fiber.Ctx) error {
	id := c.Params("id")

	ctx := c.Context()
	idInt, err := strconv.Atoi(id) // Change this line
	if err != nil {
		response := dto.ErrorResponse{
			StatusCode: 400,
			Message:    "Bad Request, " + err.Error(),
			Error:      nil,
		}
		return c.Status(400).JSON(response)
	}

	var request dto.UserGroupRequest
	err = c.BodyParser(&request)
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

	userGroupResponse, err := controller.userGroupService.UpdateUserGroup(ctx, idInt, &request)
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
		Data:       userGroupResponse,
	}

	return c.Status(200).JSON(response)
}

func (controller *UserGroupController) DeleteUserGroup(c *fiber.Ctx) error {
	id := c.Params("id")

	ctx := c.Context()
	idInt, err := strconv.Atoi(id) // Change this line
	if err != nil {
		response := dto.ErrorResponse{
			StatusCode: 400,
			Message:    "Bad Request, " + err.Error(),
			Error:      nil,
		}
		return c.Status(400).JSON(response)
	}

	userGroupResponse, err := controller.userGroupService.DeleteUserGroup(ctx, idInt)
	if err != nil {
		status_code, message := utils.GetStatusCodeFromError(err)
		if err.Error() == "cannot delete this user group" {
			status_code = http.StatusForbidden
			message = "Forbidden"
		}
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
		Data:       userGroupResponse,
	}

	return c.Status(200).JSON(response)
}
