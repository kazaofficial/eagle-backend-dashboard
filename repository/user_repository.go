package repository

import (
	"context"
	"eagle-backend-dashboard/entity"
	"strings"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) GetUser(ctx context.Context, limit *int, offset *int, sort *string, search string) ([]entity.User, error) {
	var users []entity.User
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
	if search != "" {
		// query where name like %search% or username like %search%
		query = query.Where("LOWER(name) LIKE ? OR LOWER(username) LIKE ?", "%"+strings.ToLower(search)+"%", "%"+strings.ToLower(search)+"%")
	}
	query = query.Where("id != ?", 1).Preload("UserGroup").Preload("CreatedByUser")
	err := query.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

func (r *UserRepositoryImpl) CountUser(ctx context.Context, search string) (int, error) {
	var count int64
	query := r.db.Model(&entity.User{})
	if search != "" {
		// query where name like %search% or username like %search%
		query = query.Where("LOWER(name) LIKE ? OR LOWER(username) LIKE ?", "%"+strings.ToLower(search)+"%", "%"+strings.ToLower(search)+"%")
	}
	err := query.Where("id != ?", 1).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (r *UserRepositoryImpl) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return &user, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) GetUserByID(ctx context.Context, id int, me bool) (*entity.User, error) {
	var user entity.User
	query := r.db
	if !me {
		query = query.Where("id != ?", 1)
	}
	err := query.Where("id = ?", id).Preload("UserGroup").First(&user).Error
	if err != nil {
		return &user, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) CreateUser(ctx context.Context, user *entity.User) error {
	err := r.db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) UpdateUser(ctx context.Context, user *entity.User) error {
	err := r.db.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) DeleteUser(ctx context.Context, id int) error {
	var user entity.User
	err := r.db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}
