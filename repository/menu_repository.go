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

func (r *MenuRepositoryImpl) GetMenuByIDAndUserGroupID(ctx context.Context, id int, userGroupID int) (*entity.Menu, error) {
	var menu entity.Menu
	// menu join to user_group_menu
	err := r.db.
		Joins("JOIN user_group_menus ON menus.id = user_group_menus.menu_id").
		Where("user_group_menus.user_group_id = ?", userGroupID).
		Where("menus.id = ?", id).
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
