package controller

import (
	"eagle-backend-dashboard/dto"
	"eagle-backend-dashboard/service"
	"eagle-backend-dashboard/utils"
	"net/http"

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
