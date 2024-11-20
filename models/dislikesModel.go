package models

import "time"

type Dislikes struct {
	VideoId    uint      `json:"videoId" gorm:"uniqueIndex:idx_video_user"`
	UserId     uint      `json:"userId" gorm:"uniqueIndex:idx_video_user"`
	Video      Video     `json:"video" gorm:"foreignkey:VideoId"`
	User       User      `json:"user" gorm:"foreignkey:UserId"`
	DislikedAt time.Time `json:"Dislike_timeStamp"`
}

func (l *Dislikes) BeforeCreate() (err error) {
	l.DislikedAt = GetCurrentTime()
	return
}

func (l *Dislikes) BeforeUpdate() (err error) {
	l.DislikedAt = GetCurrentTime()
	return
}