package models

import (
	"strings"

	"gorm.io/gorm"
)

type Video struct {
	User              User   `json:"user" gorm:"foreignkey:UserID"`
	UserID            uint   `json:"user_id"` // Foreign key
	ID                uint   `json:"id" gorm:"primary_key auto_increment"`
	UUID              string `json:"uuid" gorm:"type:varchar(100);unique;not null"`
	Name              string `json:"name" gorm:"type:varchar(100);not null"`
	Tags              string `json:"tags" gorm:"type:varchar(100);not null"`
	Description       string `json:"description" gorm:"type:varchar(100);not null"`
	Path              string `json:"path" gorm:"type:varchar(100);not null"`
	OriginalVideoPath string `json:"original_video_path" gorm:"type:varchar(100);not null"`
}

func (v *Video) AfterCreate(tx *gorm.DB) (err error) {
	tags := strings.Split(v.Tags, ",")

	for _, tag := range tags {
		tag = strings.TrimSpace(tag)
		trendingTag := TrendingTags{
			Tag:         tag,
			Usage_count: 0,
		}

		if err := tx.FirstOrCreate(&trendingTag, TrendingTags{Tag: tag}).Error; err != nil {
			return err
		}
	}

	return nil
}
