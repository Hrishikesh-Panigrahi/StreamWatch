package models

type Video struct {
	User   User   `json:"user" gorm:"foreignkey:UserID"`
	UserID uint   `json:"user_id"` // Foreign key
	ID     uint   `json:"id" gorm:"primary_key auto_increment"`
	UUID   string `json:"uuid" gorm:"type:varchar(100);unique;not null"`
	Name   string `json:"name" gorm:"type:varchar(100);not null"`
	Path   string `json:"path" gorm:"type:varchar(100);not null"`
}
