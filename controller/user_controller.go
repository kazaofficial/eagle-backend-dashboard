package controller

import (
	"eagle-backend-dashboard/dto"
	"eagle-backend-dashboard/service"
	"eagle-backend-dashboard/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService service.UserService
}

func NewUserRoutes(handler fiber.Router, userService service.UserService) {
	r := &UserController{
		userService: userService,
	}

	handler.Get("/user", r.GetUser)
	handler.Get("/user/:id", r.GetUserByID)
	handler.Post("/user", r.CreateUser)
	handler.Put("/user/:id", r.UpdateUser)
	handler.Delete("/user/:id", r.DeleteUser)

	handler.Get("/me", r.GetMe)
	handler.Put("/me", r.UpdateMe)
}

func (controller *UserController) GetUser(c *fiber.Ctx) error {
	var request dto.UserListRequest

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

	users, pagination, err := controller.userService.GetUser(ctx, &request)
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
			Data:       users,
		},
		Pagination: pagination,
	}
	return c.Status(200).JSON(response)
}

func (controller *UserController) GetUserByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		response := dto.ErrorResponse{
			StatusCode: 400,
			Message:    "Bad Request, " + err.Error(),
			Error:      nil,
		}
		return c.Status(400).JSON(response)
	}

	ctx := c.Context()
	user, err := controller.userService.GetUserByID(ctx, id, false)
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
		Data:       user,
	}
	return c.Status(200).JSON(response)
}

func (controller *UserController) CreateUser(c *fiber.Ctx) error {
	var request dto.UserRequest

	ctx := c.Context()
	err := c.BodyParser(&request)
	userID := c.Locals("id").(int)

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

	user, err := controller.userService.CreateUser(ctx, userID, &request)
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
		Data:       user,
	}
	return c.Status(200).JSON(response)
}

func (controller *UserController) UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		response := dto.ErrorResponse{
			StatusCode: 400,
			Message:    "Bad Request, " + err.Error(),
			Error:      nil,
		}
		return c.Status(400).JSON(response)
	}

	var request dto.UserUpdateRequest

	ctx := c.Context()
	userID := c.Locals("id").(int)
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

	user, err := controller.userService.UpdateUser(ctx, id, false, userID, &request)
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
		Data:       user,
	}
	return c.Status(200).JSON(response)
}

func (controller *UserController) DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		response := dto.ErrorResponse{
			StatusCode: 400,
			Message:    "Bad Request, " + err.Error(),
			Error:      nil,
		}
		return c.Status(400).JSON(response)
	}

	ctx := c.Context()
	user, err := controller.userService.DeleteUser(ctx, id)
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
		Data:       user,
	}
	return c.Status(200).JSON(response)
}

func (controller *UserController) GetMe(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Locals("id").(int)
	user, err := controller.userService.GetUserByID(ctx, id, true)
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
		Data:       user,
	}
	return c.Status(200).JSON(response)
}

func (controller *UserController) UpdateMe(c *fiber.Ctx) error {
	id := c.Locals("id").(int)

	var request dto.UserUpdateRequest

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

	user, err := controller.userService.UpdateUser(ctx, id, true, id, &request)
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
		Data:       user,
	}
	return c.Status(200).JSON(response)
}
