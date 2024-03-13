// CREATE TABLE IF NOT EXISTS users (
//     id SERIAL PRIMARY KEY,
//     group_id INT NOT NULL REFERENCES user_groups(id) ON DELETE SET NULL ON UPDATE CASCADE,
//     name VARCHAR(100) NOT NULL,
//     usename VARCHAR(100) NOT NULL UNIQUE,
//     password TEXT NOT NULL,
//     nrp VARCHAR(100) NOT NULL,
//     created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
//     updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
//     deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
// );

package entity

import (
	"time"
)

type User struct {
	ID        int        `gorm:"primary_key" json:"id"`
	GroupID   int        `json:"group_id"`
	Name      string     `json:"name"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	NRP       string     `json:"nrp"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
