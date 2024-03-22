package service

import (
	"context"
	"eagle-backend-dashboard/dto"
	"eagle-backend-dashboard/entity"
	"eagle-backend-dashboard/repository"
	"fmt"
)

type userGroupMenuServiceImpl struct {
	userGroupMenuRepository repository.UserGroupMenuRepository
}

// NewUserGroupMenuService creates a new instance of UserGroupMenuService.
func NewUserGroupMenuService(userGroupMenuRepository repository.UserGroupMenuRepository) UserGroupMenuService {
	return &userGroupMenuServiceImpl{
		userGroupMenuRepository: userGroupMenuRepository,
	}
}

func (s *userGroupMenuServiceImpl) GetUserGroupMenuByParentIDAndUserID(ctx context.Context, parentID int, userID int) ([]dto.UserGroupMenuResponse, error) {
	userGroupMenus, err := s.userGroupMenuRepository.GetUserGroupMenuByParentIDAndUserID(ctx, parentID, userID)
	if err != nil {
		return nil, err
	}

	userGroupMenuResponses := []dto.UserGroupMenuResponse{}
	for _, userGroupMenu := range userGroupMenus {
		userGroupMenuResponses = append(userGroupMenuResponses, ConvertUserGroupMenuResponseFromEntity(userGroupMenu))
	}

	return userGroupMenuResponses, nil
}

func (s *userGroupMenuServiceImpl) CreateManyUserGroupMenu(ctx context.Context, request dto.UserGroupMenuRequest) error {
	userGroupMenus := []entity.UserGroupMenu{}
	for _, menuID := range request.MenuIDs {
		// check if user group menu already exists
		userGroupMenu, _ := s.userGroupMenuRepository.GetByUserGroupIDAndMenuID(ctx, request.UserGroupID, menuID)
		if userGroupMenu == nil {
			userGroupMenu := entity.UserGroupMenu{
				UserGroupID: request.UserGroupID,
				MenuID:      menuID,
			}
			userGroupMenus = append(userGroupMenus, userGroupMenu)
		}
	}

	if len(userGroupMenus) == 0 {
		return fmt.Errorf("All user group menu already exists")
	}
	return s.userGroupMenuRepository.CreateManyUserGroupMenu(ctx, userGroupMenus)
}

func (s *userGroupMenuServiceImpl) DeleteManyUserGroupMenu(ctx context.Context, request dto.UserGroupMenuRequest) error {
	return s.userGroupMenuRepository.DeleteManyUserGroupMenuByUserGroupIDAndMenuIDs(ctx, request.UserGroupID, request.MenuIDs)
}

func ConvertUserGroupMenuResponseFromEntity(userGroupMenu entity.UserGroupMenu) dto.UserGroupMenuResponse {
	return dto.UserGroupMenuResponse{
		ID:          userGroupMenu.ID,
		UserGroupID: userGroupMenu.UserGroupID,
		MenuID:      userGroupMenu.MenuID,
		CreatedAt:   userGroupMenu.CreatedAt,
		UpdatedAt:   userGroupMenu.UpdatedAt,
	}
}
