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

type ManajemenDataProsesController struct {
	manajemenDataProsesService service.ManajemenDataProsesService
}

func NewManajemenDataProsesRoutes(handler fiber.Router, manajemenDataProsesService service.ManajemenDataProsesService) {
	r := &ManajemenDataProsesController{
		manajemenDataProsesService: manajemenDataProsesService,
	}

	handler.Get("/manajemen-data-proses/ssh", r.TestSSHToServer)
	handler.Get("manajemen-data-proses/penarikan-data", r.GetDaftarProsesPenarikanData)
	handler.Get("manajemen-data-proses/penarikan-data/:id", r.GetDaftarProsesPenarikanDataByID)
	handler.Post("manajemen-data-proses/penarikan-data", r.CreateDaftarProsesPenarikanData)
	handler.Put("manajemen-data-proses/penarikan-data/:id", r.UpdateDaftarProsesPenarikanData)
	handler.Delete("manajemen-data-proses/penarikan-data/:id", r.DeleteDaftarProsesPenarikanData)
}

func (r *ManajemenDataProsesController) TestSSHToServer(c *fiber.Ctx) error {
	err := r.manajemenDataProsesService.TestSSHToServer()
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
		Data:       interface{}(nil),
	}

	return c.Status(200).JSON(response)
}

func (controller *ManajemenDataProsesController) GetDaftarProsesPenarikanData(c *fiber.Ctx) error {
	var request dto.DaftarProsesPenarikanDataListRequest

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

	daftarProsesPenarikanDataResponses, pagination, err := controller.manajemenDataProsesService.GetDaftarProsesPenarikanData(ctx, &request)
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
			Data:       daftarProsesPenarikanDataResponses,
		},
		Pagination: pagination,
	}

	return c.Status(200).JSON(response)
}

func (controller *ManajemenDataProsesController) GetDaftarProsesPenarikanDataByID(c *fiber.Ctx) error {
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
	daftarProsesPenarikanDataResponse, err := controller.manajemenDataProsesService.GetDaftarProsesPenarikanDataByID(ctx, idInt)
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
		Data:       daftarProsesPenarikanDataResponse,
	}

	return c.Status(200).JSON(response)
}

func (controller *ManajemenDataProsesController) CreateDaftarProsesPenarikanData(c *fiber.Ctx) error {
	var request dto.DaftarProsesPenarikanDataRequest

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

	daftarProsesPenarikanDataResponse, err := controller.manajemenDataProsesService.CreateDaftarProsesPenarikanData(ctx, &request)
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
		Data:       daftarProsesPenarikanDataResponse,
	}

	return c.Status(200).JSON(response)
}

func (controller *ManajemenDataProsesController) UpdateDaftarProsesPenarikanData(c *fiber.Ctx) error {
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

	var request dto.DaftarProsesPenarikanDataRequest
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

	daftarProsesPenarikanDataResponse, err := controller.manajemenDataProsesService.UpdateDaftarProsesPenarikanData(ctx, idInt, &request)
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
		Data:       daftarProsesPenarikanDataResponse,
	}

	return c.Status(200).JSON(response)
}

func (controller *ManajemenDataProsesController) DeleteDaftarProsesPenarikanData(c *fiber.Ctx) error {
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

	daftarProsesPenarikanDataResponse, err := controller.manajemenDataProsesService.DeleteDaftarProsesPenarikanData(ctx, idInt)
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
		Data:       daftarProsesPenarikanDataResponse,
	}

	return c.Status(200).JSON(response)
}
