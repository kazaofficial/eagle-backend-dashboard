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

func (service *MenuServiceImpl) GetMenuByUrlKeyAndUserGroupID(ctx context.Context, urlKey string, userGroupID int) (*dto.MenuResponse, error) {
	menu, err := service.menuRepository.GetMenuByUrlKeyAndUserGroupID(ctx, urlKey, userGroupID)
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
		UrlKey:      entity.UrlKey,
		Description: entity.Description,
		Icon:        entity.Icon,
		Url:         entity.Url,
		SubMenus:    subMenus,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}
