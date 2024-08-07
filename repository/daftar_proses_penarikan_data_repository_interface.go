package repository

import (
	"context"
	"eagle-backend-dashboard/entity"
)

// DaftarProsesPenarikanDataRepository defines the methods for interacting with user groups.
type DaftarProsesPenarikanDataRepository interface {
	GetDaftarProsesPenarikanData(ctx context.Context, limit *int, offset *int, sort *string, search string) ([]entity.DaftarProsesPenarikanData, error)
	CountDaftarProsesPenarikanData(ctx context.Context, search string) (int, error)
	GetDaftarProsesPenarikanDataByID(ctx context.Context, id int) (*entity.DaftarProsesPenarikanData, error)
	CreateDaftarProsesPenarikanData(ctx context.Context, userGroup *entity.DaftarProsesPenarikanData) error
	UpdateDaftarProsesPenarikanData(ctx context.Context, userGroup *entity.DaftarProsesPenarikanData) error
	DeleteDaftarProsesPenarikanData(ctx context.Context, id int) error
}
