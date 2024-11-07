package models

import (
	"time"
)

type Likes struct {
	VideoId string    `json:"videoId" gorm:"uniqueIndex:idx_video_user"`
	UserId  string    `json:"userId" gorm:"uniqueIndex:idx_video_user"`
	Video   Video     `json:"video" gorm:"foreignkey:VideoId"`
	User    User      `json:"user" gorm:"foreignkey:UserId"`
	LikedAt time.Time `json:"timeStamp"`
}
