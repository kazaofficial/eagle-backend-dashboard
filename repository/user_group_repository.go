package repository

import (
	"context"
	"eagle-backend-dashboard/entity"
	"strings"

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

func (r *UserGroupRepositoryImpl) GetUserGroup(ctx context.Context, limit *int, offset *int, sort *string, search string) ([]entity.UserGroup, error) {
	var userGroups []entity.UserGroup
	query := r.db
	query = query.Joins("LEFT JOIN users ON user_groups.id = users.user_group_id AND users.deleted_at IS NULL").
		Select("user_groups.*, COUNT(users.id) AS number_of_users").
		Group("user_groups.id")
	if limit != nil {
		query = query.Limit(*limit)
	}
	if offset != nil {
		query = query.Offset(*offset)
	}
	if sort != nil {
		query = query.Order(*sort)
	}
	if search != "" {

		query = query.Where("LOWER(user_groups.name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	query = query.Where("user_groups.id != ?", 1)
	err := query.Find(&userGroups).Error
	if err != nil {
		return nil, err
	}
	return userGroups, nil
}

func (r *UserGroupRepositoryImpl) CountUserGroup(ctx context.Context, search string) (int, error) {
	var count int64
	query := r.db.Model(&entity.UserGroup{}).Where("user_groups.id != ?", 1)
	if search != "" {
		// convert to lower case at where
		query = query.Where("LOWER(user_groups.name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	err := query.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (r *UserGroupRepositoryImpl) GetUserGroupByID(ctx context.Context, id int) (*entity.UserGroup, error) {
	var userGroup entity.UserGroup
	err := r.db.Where("id = ?", id).Where("user_groups.id != ?", 1).First(&userGroup).Error
	if err != nil {
		return nil, err
	}
	return &userGroup, nil
}

func (r *UserGroupRepositoryImpl) CreateUserGroup(ctx context.Context, userGroup *entity.UserGroup) error {
	err := r.db.Create(userGroup).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserGroupRepositoryImpl) UpdateUserGroup(ctx context.Context, userGroup *entity.UserGroup) error {
	err := r.db.Save(userGroup).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserGroupRepositoryImpl) DeleteUserGroup(ctx context.Context, id int) error {
	err := r.db.Delete(&entity.UserGroup{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
