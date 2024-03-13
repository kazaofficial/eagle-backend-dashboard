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

func (r *MenuRepositoryImpl) GetMenu(ctx context.Context, limit *int, offset *int, sort *string) ([]entity.Menu, error) {
	var menus []entity.Menu
	query := r.db
	if limit != nil {
		query = query.Limit(*limit)
	}
	if offset != nil {
		query = query.Offset(*offset)
	}
	if sort != nil {
		query = query.Order(*sort)
	}
	query = query.Where("parent_id = ?", 1)
	query = query.Preload("SubMenus.SubMenus.SubMenus").Find(&menus)
	err := query.Find(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (r *MenuRepositoryImpl) CountMenu(ctx context.Context) (int, error) {
	var count int64
	err := r.db.Model(&entity.Menu{}).Where("parent_id = ?", 1).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
