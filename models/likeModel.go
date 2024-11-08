package models

import (
	"time"
)

type Likes struct {
	VideoId uint      `json:"videoId" gorm:"uniqueIndex:idx_video_user"`
	UserId  uint      `json:"userId" gorm:"uniqueIndex:idx_video_user"`
	Video   Video     `json:"video" gorm:"foreignkey:VideoId"`
	User    User      `json:"user" gorm:"foreignkey:UserId"`
	LikedAt time.Time `json:"timeStamp"`
}

func GetCurrentTime() time.Time {

	return time.Now()

}
