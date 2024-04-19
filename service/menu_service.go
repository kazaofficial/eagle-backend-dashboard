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
		menuResponses = append(menuResponses, ConvertMenuEntityToDTO(menu))
	}

	return menuResponses, nil
}

func (service *MenuServiceImpl) GetMainMenu(ctx context.Context) ([]dto.MenuResponse, error) {
	menus, err := service.menuRepository.GetMainMenu(ctx)
	if err != nil {
		return nil, err
	}

	menuResponses := []dto.MenuResponse{}
	for _, menu := range menus {
		menuResponses = append(menuResponses, ConvertMenuEntityToDTO(menu))
	}

	return menuResponses, nil
}

func (service *MenuServiceImpl) GetMenuByUrlKeyAndUserGroupID(ctx context.Context, urlKey string, userGroupID int) (*dto.MenuResponse, error) {
	menu, err := service.menuRepository.GetMenuByUrlKeyAndUserGroupID(ctx, urlKey, userGroupID)
	if err != nil {
		return nil, err
	}

	menuResponse := ConvertMenuWithSubMenusEntityToDTO(*menu)
	return &menuResponse, nil
}

func ConvertMenuWithSubMenusEntityToDTO(entity entity.MenuWithSubMenus) dto.MenuResponse {
	subMenus := []dto.MenuResponse{}
	for _, subMenu := range entity.SubMenus {
		subMenus = append(subMenus, ConvertMenuWithSubMenusEntityToDTO(subMenu))
	}

	newMenu := ConvertMenuEntityToDTO(entity.Menu)
	newMenu.SubMenus = subMenus
	return newMenu
}

func ConvertMenuEntityToDTO(entity entity.Menu) dto.MenuResponse {
	return dto.MenuResponse{
		ID:          entity.ID,
		Name:        entity.Name,
		ParentID:    entity.ParentID,
		UrlKey:      entity.UrlKey,
		Description: entity.Description,
		Icon:        entity.Icon,
		Url:         entity.Url,
		IsShown:     entity.IsShown,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}
