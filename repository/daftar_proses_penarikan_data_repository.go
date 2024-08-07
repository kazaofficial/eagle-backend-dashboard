package repository

import (
	"context"
	"eagle-backend-dashboard/entity"
	"strings"

	"gorm.io/gorm"
)

type DaftarProsesPenarikanDataRepositoryImpl struct {
	db *gorm.DB
}

func NewDaftarProsesPenarikanDataRepository(db *gorm.DB) DaftarProsesPenarikanDataRepository {
	return &DaftarProsesPenarikanDataRepositoryImpl{
		db: db,
	}
}

func (r *DaftarProsesPenarikanDataRepositoryImpl) GetDaftarProsesPenarikanData(ctx context.Context, limit *int, offset *int, sort *string, search string) ([]entity.DaftarProsesPenarikanData, error) {
	var DaftarProsesPenarikanData []entity.DaftarProsesPenarikanData
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
		query = query.Where("(LOWER(daftar_proses_penarikan_data.database_type) LIKE ? OR LOWER(daftar_proses_penarikan_data.source_connection_id) LIKE ?)", "%"+strings.ToLower(search)+"%", "%"+strings.ToLower(search)+"%")
	}
	err := query.Find(&DaftarProsesPenarikanData).Error
	if err != nil {
		return nil, err
	}
	return DaftarProsesPenarikanData, nil
}

func (r *DaftarProsesPenarikanDataRepositoryImpl) CountDaftarProsesPenarikanData(ctx context.Context, search string) (int, error) {
	var count int64
	query := r.db.Model(&entity.DaftarProsesPenarikanData{})
	if search != "" {
		// convert to lower case at where
		query.Where("(LOWER(daftar_proses_penarikan_data.database_type) LIKE ? OR LOWER(daftar_proses_penarikan_data.source_connection_id) LIKE ?)", "%"+strings.ToLower(search)+"%", "%"+strings.ToLower(search)+"%")
	}
	err := query.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (r *DaftarProsesPenarikanDataRepositoryImpl) GetDaftarProsesPenarikanDataByID(ctx context.Context, id int) (*entity.DaftarProsesPenarikanData, error) {
	var daftarProsesPenarikanData entity.DaftarProsesPenarikanData
	err := r.db.Where("id = ?", id).First(&daftarProsesPenarikanData).Error
	if err != nil {
		return nil, err
	}
	return &daftarProsesPenarikanData, nil
}

func (r *DaftarProsesPenarikanDataRepositoryImpl) CreateDaftarProsesPenarikanData(ctx context.Context, daftarProsesPenarikanData *entity.DaftarProsesPenarikanData) error {
	err := r.db.Create(daftarProsesPenarikanData).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *DaftarProsesPenarikanDataRepositoryImpl) UpdateDaftarProsesPenarikanData(ctx context.Context, daftarProsesPenarikanData *entity.DaftarProsesPenarikanData) error {
	err := r.db.Save(daftarProsesPenarikanData).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *DaftarProsesPenarikanDataRepositoryImpl) DeleteDaftarProsesPenarikanData(ctx context.Context, id int) error {
	err := r.db.Delete(&entity.DaftarProsesPenarikanData{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
