package service

import (
	"context"
	"eagle-backend-dashboard/dto"
	"eagle-backend-dashboard/entity"
	"eagle-backend-dashboard/repository"
	"strings"
)

type MenuServiceImpl struct {
	menuRepository repository.MenuRepository
}

func NewMenuService(menuRepository repository.MenuRepository) MenuService {
	return &MenuServiceImpl{
		menuRepository: menuRepository,
	}
}

func (s *MenuServiceImpl) GetMenu(ctx context.Context, request *dto.MenuRequest) ([]dto.MenuResponse, *dto.Pagination, error) {
	offset := 0
	page := 1
	limit := 10

	if request.Page != nil {
		page = *request.Page
	}

	if request.Limit != nil {
		limit = *request.Limit
	}

	if page > 1 {
		offset = (page - 1) * limit
	}

	sort := "id desc"
	if request.Sort != "" {
		sort = request.Sort
		sort = strings.ReplaceAll(sort, ".", " ")
	}

	menus, err := s.menuRepository.GetMenu(ctx, &limit, &offset, &sort)
	if err != nil {
		return nil, nil, err
	}
	count, err := s.menuRepository.CountMenu(ctx)
	if err != nil {
		return nil, nil, err
	}

	menuResponses := []dto.MenuResponse{}
	for _, menu := range menus {
		menuResponses = append(menuResponses, ConverMenuEntityToDTO(menu))
	}

	pagination := dto.Pagination{
		Page:      page,
		Limit:     limit,
		Total:     len(menus),
		TotalData: count,
		TotalPage: count/limit + 1,
	}

	return menuResponses, &pagination, nil
}

func ConverMenuEntityToDTO(entity entity.Menu) dto.MenuResponse {
	return dto.MenuResponse{
		ID:          entity.ID,
		Name:        entity.Name,
		ParentID:    entity.ParentID,
		Description: entity.Description,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}
