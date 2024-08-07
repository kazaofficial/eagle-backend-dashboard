package service

import (
	"context"
	"eagle-backend-dashboard/dto"
)

// MenuService defines the methods for interacting with menus.
type ManajemenDataProsesService interface {
	TestSSHToServer() error
	GetDaftarProsesPenarikanData(ctx context.Context, request *dto.DaftarProsesPenarikanDataListRequest) ([]dto.DaftarProsesPenarikanDataResponse, *dto.Pagination, error)
	GetDaftarProsesPenarikanDataByID(ctx context.Context, id int) (*dto.DaftarProsesPenarikanDataResponse, error)
	CreateDaftarProsesPenarikanData(ctx context.Context, request *dto.DaftarProsesPenarikanDataRequest) (*dto.DaftarProsesPenarikanDataResponse, error)
	UpdateDaftarProsesPenarikanData(ctx context.Context, id int, request *dto.DaftarProsesPenarikanDataRequest) (*dto.DaftarProsesPenarikanDataResponse, error)
	DeleteDaftarProsesPenarikanData(ctx context.Context, id int) (*dto.DaftarProsesPenarikanDataResponse, error)
}
