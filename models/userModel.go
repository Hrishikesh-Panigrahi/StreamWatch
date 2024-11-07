package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID          uint         `json:"id" gorm:"primary_key"`
	Name        string       `json:"name" gorm:"type:varchar(100);not null"`
	Email       string       `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password    string       `json:"password" gorm:"type:varchar(100);not null"`
	Age         uint8        `json:"age" gorm:"type:tinyint;not null"`
	Birthday    *time.Time   `json:"birthday" gorm:"type:date"`
	Is_verified bool         `json:"is_verified" gorm:"type:boolean;default:false"`
	ActivatedAt sql.NullTime `json:"activated_at" gorm:"type:timestamp"`
	CreatedAt   time.Time    `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt   time.Time    `json:"updated_at" gorm:"type:timestamp"`
}
