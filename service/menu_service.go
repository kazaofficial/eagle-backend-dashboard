package service

import (
	"context"
	"eagle-backend-dashboard/dto"
	"eagle-backend-dashboard/entity"
	"eagle-backend-dashboard/repository"
)

type MenuServiceImpl struct {
	menuRepository repository.MenuRepository
}

func NewMenuService(menuRepository repository.MenuRepository) MenuService {
	return &MenuServiceImpl{
		menuRepository: menuRepository,
	}
}

func (service *MenuServiceImpl) GetMenuByUserGroupID(ctx context.Context, userGroupID int) ([]dto.MenuResponse, error) {
	menus, err := service.menuRepository.GetMenuByUserGroupID(ctx, userGroupID)
	if err != nil {
		return nil, err
	}

	menuResponses := []dto.MenuResponse{}
	for _, menu := range menus {
		menuResponses = append(menuResponses, ConverMenuEntityToDTO(menu))
	}

	return menuResponses, nil
}

func (service *MenuServiceImpl) GetMenuByIDAndUserGroupID(ctx context.Context, id int, userGroupID int) (*dto.MenuResponse, error) {
	menu, err := service.menuRepository.GetMenuByIDAndUserGroupID(ctx, id, userGroupID)
	if err != nil {
		return nil, err
	}

	menuResponse := ConverMenuEntityToDTO(*menu)
	return &menuResponse, nil
}

func ConverMenuEntityToDTO(entity entity.Menu) dto.MenuResponse {
	subMenus := []dto.MenuResponse{}
	for _, subMenu := range entity.SubMenus {
		subMenus = append(subMenus, ConverMenuEntityToDTO(subMenu))
	}
	return dto.MenuResponse{
		ID:          entity.ID,
		Name:        entity.Name,
		ParentID:    entity.ParentID,
		Description: entity.Description,
		SubMenus:    subMenus,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}
