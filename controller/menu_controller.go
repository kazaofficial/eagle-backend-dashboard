package controller

import (
	"eagle-backend-dashboard/dto"
	"eagle-backend-dashboard/service"
	"eagle-backend-dashboard/utils"

	"github.com/gofiber/fiber/v2"
)

type MenuController struct {
	menuService service.MenuService
}

func NewMenuRoutes(handler fiber.Router, menuService service.MenuService) {
	r := &MenuController{
		menuService: menuService,
	}

	handler.Get("/menu", r.GetMenuByUserGroupID)
	handler.Get("/menu/:id", r.GetMenuByIDAndUserGroupID)
}

func (r *MenuController) GetMenuByUserGroupID(c *fiber.Ctx) error {
	userGroupID := c.Locals("user_group_id").(int)
	menuResponses, err := r.menuService.GetMenuByUserGroupID(c.Context(), userGroupID)
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
		Data:       menuResponses,
	}

	return c.Status(200).JSON(response)
}

func (r *MenuController) GetMenuByIDAndUserGroupID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		response := dto.ErrorResponse{
			StatusCode: 400,
			Message:    "Bad Request, " + err.Error(),
			Error:      nil,
		}
		return c.Status(400).JSON(response)
	}
	userGroupID := c.Locals("user_group_id").(int)
	menuResponse, err := r.menuService.GetMenuByIDAndUserGroupID(c.Context(), id, userGroupID)
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
