package service

import (
	"context"
	"eagle-backend-dashboard/dto"
	"eagle-backend-dashboard/entity"
	"eagle-backend-dashboard/repository"
	"errors"
	"strings"
)

// UserGroupServiceImpl implements the UserGroupService interface.
type UserGroupServiceImpl struct {
	userGroupRepository repository.UserGroupRepository
	menuRepository      repository.MenuRepository
}

// NewUserGroupService creates a new UserGroupService.
func NewUserGroupService(userGroupRepository repository.UserGroupRepository, menuRepository repository.MenuRepository) UserGroupService {
	return &UserGroupServiceImpl{
		userGroupRepository: userGroupRepository,
		menuRepository:      menuRepository,
	}
}

func (service *UserGroupServiceImpl) GetUserGroup(ctx context.Context, request *dto.UserGroupListRequest) ([]dto.UserGroupResponse, *dto.Pagination, error) {
	offset := 0
	page := 1
	limit := 10

	if request.Page != nil {
		page = *request.Page
		offset = (page - 1) * limit
	}

	if request.Limit != nil {
		limit = *request.Limit
	}

	sort := "id desc"
	if request.Sort != "" {
		sort = request.Sort
		sort = strings.ReplaceAll(sort, ".", " ")
	}

	userGroups, err := service.userGroupRepository.GetUserGroup(ctx, &limit, &offset, &sort)
	if err != nil {
		return nil, nil, err
	}

	countUserGroups, err := service.userGroupRepository.CountUserGroup(ctx)
	if err != nil {
		return nil, nil, err
	}

	userGroupResponses := []dto.UserGroupResponse{}
	for _, userGroup := range userGroups {
		userGroupResponses = append(userGroupResponses, ConvertUserGroupEntityToDTO(userGroup))
	}

	pagination := dto.Pagination{
		Page:      page,
		Limit:     limit,
		Total:     len(userGroups),
		TotalData: countUserGroups,
		TotalPage: countUserGroups/limit + 1,
	}

	return userGroupResponses, &pagination, nil
}

func (service *UserGroupServiceImpl) GetUserGroupWithAccess(ctx context.Context, request *dto.UserGroupListRequest) ([]dto.UserGroupResponse, *dto.Pagination, error) {
	offset := 0
	page := 1
	limit := 10

	if request.Page != nil {
		page = *request.Page
		offset = (page - 1) * limit
	}

	if request.Limit != nil {
		limit = *request.Limit
	}

	sort := "id desc"
	if request.Sort != "" {
		sort = request.Sort
		sort = strings.ReplaceAll(sort, ".", " ")
	}

	userGroups, err := service.userGroupRepository.GetUserGroup(ctx, &limit, &offset, &sort)
	if err != nil {
		return nil, nil, err
	}

	countUserGroups, err := service.userGroupRepository.CountUserGroup(ctx)
	if err != nil {
		return nil, nil, err
	}

	userGroupResponses := []dto.UserGroupResponse{}
	for _, userGroup := range userGroups {
		menus, err := service.menuRepository.GetMenuAccessByUserGroupID(ctx, userGroup.ID)
		if err == nil {
			userGroup.Menus = menus
		}
		userGroupResponses = append(userGroupResponses, ConvertUserGroupEntityToDTO(userGroup))
	}

	pagination := dto.Pagination{
		Page:      page,
		Limit:     limit,
		Total:     len(userGroups),
		TotalData: countUserGroups,
		TotalPage: countUserGroups/limit + 1,
	}

	return userGroupResponses, &pagination, nil
}

func (service *UserGroupServiceImpl) GetUserGroupByID(ctx context.Context, id int) (*dto.UserGroupResponse, error) {
	userGroup, err := service.userGroupRepository.GetUserGroupByID(ctx, id)
	if err != nil {
		return nil, err
	}

	userGroupResponse := ConvertUserGroupEntityToDTO(*userGroup)
	return &userGroupResponse, nil
}

func (service *UserGroupServiceImpl) CreateUserGroup(ctx context.Context, request *dto.UserGroupRequest) (*dto.UserGroupResponse, error) {
	userGroup := entity.UserGroup{
		Name:        request.Name,
		Description: request.Description,
	}

	err := service.userGroupRepository.CreateUserGroup(ctx, &userGroup)
	if err != nil {
		return nil, err
	}

	userGroupResponse := ConvertUserGroupEntityToDTO(userGroup)
	return &userGroupResponse, nil
}

func (service *UserGroupServiceImpl) UpdateUserGroup(ctx context.Context, id int, request *dto.UserGroupRequest) (*dto.UserGroupResponse, error) {
	userGroup, err := service.userGroupRepository.GetUserGroupByID(ctx, id)
	if err != nil {
		return nil, err
	}

	userGroup.Name = request.Name
	userGroup.Description = request.Description

	err = service.userGroupRepository.UpdateUserGroup(ctx, userGroup)
	if err != nil {
		return nil, err
	}

	userGroupResponse := ConvertUserGroupEntityToDTO(*userGroup)
	return &userGroupResponse, nil
}

func (service *UserGroupServiceImpl) DeleteUserGroup(ctx context.Context, id int) (*dto.UserGroupResponse, error) {
	userGroup, err := service.userGroupRepository.GetUserGroupByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if userGroup.ID == 1 {
		return nil, errors.New("cannot delete this user group")
	}

	err = service.userGroupRepository.DeleteUserGroup(ctx, id)
	if err != nil {
		return nil, err
	}

	userGroupResponse := ConvertUserGroupEntityToDTO(*userGroup)
	return &userGroupResponse, nil
}

func ConvertUserGroupEntityToDTO(userGroup entity.UserGroup) dto.UserGroupResponse {
	userGroupResponse := dto.UserGroupResponse{
		ID:            userGroup.ID,
		Name:          userGroup.Name,
		Description:   userGroup.Description,
		NumberOfUsers: userGroup.NumberOfUsers,
		CreatedAt:     userGroup.CreatedAt,
		UpdatedAt:     userGroup.UpdatedAt,
		DeletedAt:     userGroup.DeletedAt,
	}

	if userGroup.Menus != nil {
		menus := []dto.MenuResponse{}
		for _, menu := range userGroup.Menus {
			menus = append(menus, ConverUserGroupMenuEntityToDTO(menu))
		}
		userGroupResponse.Menus = menus
	}

	return userGroupResponse
}

func ConverUserGroupMenuEntityToDTO(entity entity.MenuWithUserGroup) dto.MenuResponse {
	subMenus := []dto.MenuResponse{}
	for _, subMenu := range entity.SubMenus {
		subMenus = append(subMenus, ConverUserGroupMenuEntityToDTO(subMenu))
	}

	isActive := false
	newMenu := dto.MenuResponse{
		ID:       entity.ID,
		Name:     entity.Name,
		ParentID: entity.ParentID,
		SubMenus: subMenus,
		IsActive: &isActive,
	}
	if entity.UserGroupID != nil {
		isActive = true
		newMenu.IsActive = &isActive
	}
	return newMenu
}
