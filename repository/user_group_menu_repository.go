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

func (r *UserGroupMenuRepositoryImpl) GetByUserGroupIDAndMenuID(ctx context.Context, userGroupID int, menuID int) (*entity.UserGroupMenu, error) {
	var userGroupMenu entity.UserGroupMenu
	err := r.db.Where("user_group_id = ? AND menu_id = ?", userGroupID, menuID).First(&userGroupMenu).Error
	if err != nil {
		return nil, err
	}
	return &userGroupMenu, nil
}

func (r *UserGroupMenuRepositoryImpl) CreateManyUserGroupMenu(ctx context.Context, userGroupMenus []entity.UserGroupMenu) error {
	return r.db.Create(&userGroupMenus).Error
}

func (r *UserGroupMenuRepositoryImpl) DeleteManyUserGroupMenuByUserGroupIDAndMenuIDs(ctx context.Context, userGroupID int, menuIDs []int) error {
	return r.db.Where("user_group_id = ? AND menu_id IN (?)", userGroupID, menuIDs).Delete(&entity.UserGroupMenu{}).Error
}
