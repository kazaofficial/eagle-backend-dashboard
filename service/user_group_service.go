package service

import (
	"context"
	"eagle-backend-dashboard/dto"
	"eagle-backend-dashboard/entity"
	"eagle-backend-dashboard/repository"
	"strings"
)

// UserGroupServiceImpl implements the UserGroupService interface.
type UserGroupServiceImpl struct {
	userGroupRepository repository.UserGroupRepository
}

// NewUserGroupService creates a new UserGroupService.
func NewUserGroupService(userGroupRepository repository.UserGroupRepository) UserGroupService {
	return &UserGroupServiceImpl{
		userGroupRepository: userGroupRepository,
	}
}

func (service *UserGroupServiceImpl) GetUserGroup(ctx context.Context, request *dto.UserGroupRequest) ([]dto.UserGroupResponse, *dto.Pagination, error) {
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
		userGroupResponses = append(userGroupResponses, ConverUserGroupEntityToDTO(userGroup))
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

	userGroupResponse := ConverUserGroupEntityToDTO(*userGroup)
	return &userGroupResponse, nil
}

func ConverUserGroupEntityToDTO(userGroup entity.UserGroup) dto.UserGroupResponse {
	userGroupResponse := dto.UserGroupResponse{
		ID:          userGroup.ID,
		Name:        userGroup.Name,
		Description: userGroup.Description,
		CreatedAt:   userGroup.CreatedAt,
		UpdatedAt:   userGroup.UpdatedAt,
		DeletedAt:   userGroup.DeletedAt,
	}

	return userGroupResponse
}
