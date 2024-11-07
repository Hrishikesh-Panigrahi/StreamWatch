package models

import (
	"time"
)

type Likes struct {
	Video   Video     `json:"video" gorm:"foreignkey:VideoId"`
	VideoId string    `json:"videoId"`
	User    User      `json:"user" gorm:"foreignkey:UserId"`
	UserId  string    `json:"userId"`
	LikedAt time.Time `json:"timeStamp"`
}
