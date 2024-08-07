package dto

import (
	"time"

	"gorm.io/gorm"
)

type DaftarProsesPenarikanDataResponse struct {
	ID                 int            `json:"id"`
	ConnType           string         `json:"conn_type"`
	SourceConnectionID string         `json:"source_connection_id"`
	TargetConnectionID string         `json:"target_connection_id"`
	SchemaTable        string         `json:"schema_table"`
	Schedule           string         `json:"schedule"`
	CreatedAt          *time.Time     `json:"created_at"`
	UpdatedAt          *time.Time     `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `json:"deleted_at"`
}

type DaftarProsesPenarikanDataRequest struct {
	ConnType           string `json:"conn_type" validate:"required"` // DatabaseType
	SourceConnectionID string `json:"source_connection_id" validate:"required"`
	TargetConnectionID string `json:"target_connection_id" validate:"required"`
	SchemaTable        string `json:"schema_table" validate:"required"`
	Schedule           string `json:"schedule" validate:"required"`
}

type DaftarProsesPenarikanDataListRequest struct {
	Page   *int   `json:"page"`
	Limit  *int   `json:"limit"`
	Sort   string `json:"sort"`
	Search string `json:"search"`
}
