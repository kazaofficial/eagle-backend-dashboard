package repository

import (
	"context"
	"eagle-backend-dashboard/entity"

	"gorm.io/gorm"
)

type UserGroupRepositoryImpl struct {
	db *gorm.DB
}

func NewUserGroupRepository(db *gorm.DB) UserGroupRepository {
	return &UserGroupRepositoryImpl{
		db: db,
	}
}

func (r *UserGroupRepositoryImpl) GetUserGroup(ctx context.Context, limit *int, offset *int, sort *string) ([]entity.UserGroup, error) {
	var userGroups []entity.UserGroup
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
	err := query.Find(&userGroups).Error
	if err != nil {
		return nil, err
	}
	return userGroups, nil
}

func (r *UserGroupRepositoryImpl) CountUserGroup(ctx context.Context) (int, error) {
	var count int64
	err := r.db.Model(&entity.UserGroup{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
