package service

import (
	"context"
	"eagle-backend-dashboard/dto"
	"eagle-backend-dashboard/entity"
	"eagle-backend-dashboard/repository"
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

func ConvertUserGroupMenuResponseFromEntity(userGroupMenu entity.UserGroupMenu) dto.UserGroupMenuResponse {
	return dto.UserGroupMenuResponse{
		ID:          userGroupMenu.ID,
		UserGroupID: userGroupMenu.UserGroupID,
		MenuID:      userGroupMenu.MenuID,
		CreatedAt:   userGroupMenu.CreatedAt,
		UpdatedAt:   userGroupMenu.UpdatedAt,
	}
}
