package models

import (
	"time"
)

type Likes struct {
	VideoId uint      `json:"videoId" gorm:"uniqueIndex:idx_video_user"`
	UserId  uint      `json:"userId" gorm:"uniqueIndex:idx_video_user"`
	Video   Video     `json:"video" gorm:"foreignkey:VideoId"`
	User    User      `json:"user" gorm:"foreignkey:UserId"`
	LikedAt time.Time `json:"Like_timeStamp"`
}

func GetCurrentTime() time.Time {
	return time.Now()
}

func (l *Likes) BeforeCreate() (err error) {
	l.LikedAt = GetCurrentTime()
	return
}

func (l *Likes) BeforeUpdate() (err error) {
	l.LikedAt = GetCurrentTime()
	return
}
