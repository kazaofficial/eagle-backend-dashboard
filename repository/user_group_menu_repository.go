package repository

import (
	"context"
	"eagle-backend-dashboard/entity"

	"gorm.io/gorm"
)

// UserGroupMenuRepository defines the methods for interacting with user group menus.
type UserGroupMenuRepositoryImpl struct {
	db *gorm.DB
}

func NewUserGroupMenuRepository(db *gorm.DB) UserGroupMenuRepository {
	return &UserGroupMenuRepositoryImpl{
		db: db,
	}
}

func (r *UserGroupMenuRepositoryImpl) GetUserGroupMenuByParentIDAndUserID(ctx context.Context, parentID int, userID int) ([]entity.UserGroupMenu, error) {
	var userGroupMenus []entity.UserGroupMenu
	// query join to menus table and get where parent_id = parentID
	err := r.db.
		Joins("JOIN menus ON user_group_menus.menu_id = menus.id").
		Where("menus.parent_id = ?", parentID).
		Find(&userGroupMenus).Error
	if err != nil {
		return nil, err
	}
	return userGroupMenus, nil
}
