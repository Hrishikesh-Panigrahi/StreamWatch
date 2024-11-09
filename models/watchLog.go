package models

import "time"

type WatchLog struct {
	VideoId        uint          `json:"videoId" gorm:"uniqueIndex:idx_video_user"`
	UserId         uint          `json:"userId" gorm:"uniqueIndex:idx_video_user"`
	Video          Video         `json:"video" gorm:"foreignkey:VideoId"`
	User           User          `json:"user" gorm:"foreignkey:UserId"`
	Watch_duration time.Duration `json:"timeStamp"`
}
