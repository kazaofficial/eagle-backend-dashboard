package repository

import (
	"context"
	"eagle-backend-dashboard/entity"

	"gorm.io/gorm"
)

type MenuRepositoryImpl struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &MenuRepositoryImpl{
		db: db,
	}
}

func (r *MenuRepositoryImpl) GetMenuByUserGroupID(ctx context.Context, userGroupID int) ([]entity.Menu, error) {
	var menus []entity.Menu
	// menu join to user_group_menu
	err := r.db.
		Joins("JOIN user_group_menus ON menus.id = user_group_menus.menu_id").
		Where("user_group_menus.user_group_id = ?", userGroupID).
		Where("menus.parent_id = ?", 1).
		Find(&menus).Error

	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (r *MenuRepositoryImpl) GetMainMenu(ctx context.Context) ([]entity.Menu, error) {
	var menus []entity.Menu
	err := r.db.
		Where("parent_id = ?", 1).
		Find(&menus).Error

	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (r *MenuRepositoryImpl) GetMenuByUrlKeyAndUserGroupID(ctx context.Context, urlKey string, userGroupID int) (*entity.MenuWithSubMenus, error) {
	var menu entity.MenuWithSubMenus
	// menu join to user_group_menu
	err := r.db.
		Joins("JOIN user_group_menus ON menus.id = user_group_menus.menu_id").
		Where("user_group_menus.user_group_id = ?", userGroupID).
		Where("menus.url_key = ?", urlKey).
		Where("menus.id != ?", 1).
		// preload submenus with join
		Preload("SubMenus", func(db *gorm.DB) *gorm.DB {
			return db.Joins("JOIN user_group_menus ON menus.id = user_group_menus.menu_id").
				Where("user_group_menus.user_group_id = ?", userGroupID)
		}).
		Preload("SubMenus.SubMenus", func(db *gorm.DB) *gorm.DB {
			return db.Joins("JOIN user_group_menus ON menus.id = user_group_menus.menu_id").
				Where("user_group_menus.user_group_id = ?", userGroupID)
		}).
		First(&menu).Error

	if err != nil {
		return nil, err
	}
	return &menu, nil
}

func (r *MenuRepositoryImpl) GetMenuAccessByUserGroupID(ctx context.Context, userGroupID int) ([]entity.MenuWithUserGroup, error) {
	var menu []entity.MenuWithUserGroup
	// menu join to user_group_menu
	err := r.db.
		Select("menus.*, user_group_menus.user_group_id as user_group_id").
		Joins("LEFT JOIN user_group_menus ON menus.id = user_group_menus.menu_id AND user_group_menus.user_group_id = ?", userGroupID).
		Where("menus.id != ?", 1).
		Where("menus.parent_id = ?", 1).
		// preload submenus with join
		Preload("SubMenus", func(db *gorm.DB) *gorm.DB {
			return db.Select("menus.*, user_group_menus.user_group_id as user_group_id").
				Joins("LEFT JOIN user_group_menus ON menus.id = user_group_menus.menu_id AND user_group_menus.user_group_id = ?", userGroupID).
				Order("menus.id ASC")
		}).
		Preload("SubMenus.SubMenus", func(db *gorm.DB) *gorm.DB {
			return db.Select("menus.*, user_group_menus.user_group_id as user_group_id").
				Joins("LEFT JOIN user_group_menus ON menus.id = user_group_menus.menu_id AND user_group_menus.user_group_id = ?", userGroupID).
				Order("menus.id ASC")
		}).
		Order("menus.id ASC").
		Find(&menu).Error

	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (r *MenuRepositoryImpl) PluckIDByIDOrParentID(ctx context.Context, ids []int) ([]int, error) {
	var result []int
	var menu entity.Menu
	err := r.db.
		Model(&menu).
		Where("id IN ? OR parent_id IN ?", ids, ids).
		Pluck("id", &result).Error

	if err != nil {
		return nil, err
	}
	return result, nil
}
