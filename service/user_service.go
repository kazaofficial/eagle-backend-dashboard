package service

import (
	"context"
	"eagle-backend-dashboard/dto"
	"eagle-backend-dashboard/entity"
	"eagle-backend-dashboard/repository"
	"math"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (service *UserServiceImpl) GetUser(ctx context.Context, request *dto.UserListRequest) ([]dto.UserResponse, *dto.Pagination, error) {
	offset := 0
	page := 1
	limit := 10
	sort := "id desc"

	if request.Limit != nil {
		limit = *request.Limit
	}

	if request.Page != nil {
		page = *request.Page
		offset = (page - 1) * limit
	}

	if request.Sort != "" {
		sort = request.Sort
		sort = strings.ReplaceAll(sort, ".", " ")
	}

	users, err := service.UserRepository.GetUser(ctx, &limit, &offset, &sort, request.Search)
	if err != nil {
		return nil, nil, err
	}

	countUsers, err := service.UserRepository.CountUser(ctx, request.Search)
	if err != nil {
		return nil, nil, err
	}

	userResponses := []dto.UserResponse{}
	for _, user := range users {
		userResponses = append(userResponses, ConvertUserEntityToDTO(user))
	}

	pagination := dto.Pagination{
		Page:      page,
		Limit:     limit,
		Total:     len(userResponses),
		TotalData: countUsers,
		TotalPage: int(math.Ceil(float64(countUsers) / float64(limit))),
	}

	return userResponses, &pagination, nil
}

func (service *UserServiceImpl) GetUserByID(ctx context.Context, id int, me bool) (*dto.UserResponse, error) {
	user, err := service.UserRepository.GetUserByID(ctx, id, me)
	if err != nil {
		return nil, err
	}
	userResponse := ConvertUserEntityToDTO(*user)
	return &userResponse, nil
}

func (service *UserServiceImpl) CreateUser(ctx context.Context, userID int, request *dto.UserRequest) (*dto.UserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := entity.User{
		UserGroupID: request.UserGroupID,
		Name:        request.Name,
		Username:    request.Username,
		Password:    string(hashedPassword),
		NRP:         request.NRP,
		CreatedBy:   userID,
		UpdatedBy:   userID,
	}
	err = service.UserRepository.CreateUser(ctx, &user)
	if err != nil {
		return nil, err
	}

	userResponse := ConvertUserEntityToDTO(user)
	return &userResponse, nil
}

func (service *UserServiceImpl) UpdateUser(ctx context.Context, id int, me bool, userID int, request *dto.UserUpdateRequest) (*dto.UserResponse, error) {
	user, err := service.UserRepository.GetUserByID(ctx, id, me)
	if err != nil {
		return nil, err
	}

	if request.Name != "" {
		user.Name = request.Name
	}

	if request.UserGroupID != 0 {
		user.UserGroupID = request.UserGroupID
	}

	if request.NRP != "" {
		user.NRP = request.NRP
	}

	user.UpdatedBy = userID

	err = service.UserRepository.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	userResponse := ConvertUserEntityToDTO(*user)
	return &userResponse, nil
}

func (service *UserServiceImpl) DeleteUser(ctx context.Context, id int) (*dto.UserResponse, error) {
	user, err := service.UserRepository.GetUserByID(ctx, id, false)
	if err != nil {
		return nil, err
	}

	err = service.UserRepository.DeleteUser(ctx, id)
	if err != nil {
		return nil, err
	}

	userResponse := ConvertUserEntityToDTO(*user)
	return &userResponse, nil
}

func ConvertUserEntityToDTO(user entity.User) dto.UserResponse {

	userResponse := dto.UserResponse{
		ID:          user.ID,
		UserGroupID: user.UserGroupID,
		Role:        user.Role,
		Name:        user.Name,
		Username:    user.Username,
		NRP:         user.NRP,
		LastLogin:   user.LastLogin,
		CreatedAt:   user.CreatedAt,
		CreatedBy:   user.CreatedBy,
		UpdatedAt:   user.UpdatedAt,
		UpdatedBy:   user.UpdatedBy,
		DeletedAt:   user.DeletedAt,
	}
	if user.UserGroup != nil {
		userGroup := ConvertUserGroupEntityToDTO(*user.UserGroup)
		userResponse.UserGroup = &userGroup
	}
	if user.CreatedByUser != nil {
		createdByUser := dto.UserResponse{
			ID:   user.CreatedByUser.ID,
			Name: user.CreatedByUser.Name,
		}
		userResponse.CreatedByUser = &createdByUser
	}
	return userResponse
}
