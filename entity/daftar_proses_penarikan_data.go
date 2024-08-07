package entity

import (
	"time"

	"gorm.io/gorm"
)

type DaftarProsesPenarikanData struct {
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

// TableName returns the table name of the DaftarProsesPenarikanData
func (u *DaftarProsesPenarikanData) TableName() string {
	return "daftar_proses_penarikan_data"
}
