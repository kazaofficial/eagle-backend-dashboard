package controller

import (
	"eagle-backend-dashboard/dto"
	"eagle-backend-dashboard/service"

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
}

func (controller *UserGroupController) GetUserGroup(c *fiber.Ctx) error {
	var request dto.UserGroupRequest

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
			Data:       userGroupResponses,
		},
		Pagination: pagination,
	}

	return c.Status(200).JSON(response)
}
