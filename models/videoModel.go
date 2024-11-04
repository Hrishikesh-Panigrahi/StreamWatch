package models

type Video struct {
	User              User   `json:"user" gorm:"foreignkey:UserID"`
	UserID            uint   `json:"user_id"` // Foreign key
	ID                uint   `json:"id" gorm:"primary_key auto_increment"`
	UUID              string `json:"uuid" gorm:"type:varchar(100);unique;not null"`
	Name              string `json:"name" gorm:"type:varchar(100);not null"`
	Path              string `json:"path" gorm:"type:varchar(100);not null"`
	OriginalVideoPath string `json:"original_video_path" gorm:"type:varchar(100);not null"`
}

// func (v *Video) AfterCreate(tx *gorm.DB) (err error) {
// 	videoPath := v.Path
// 	if len(videoPath) > 10 { // Ensure path length is sufficient before slicing
// 		videoPath = videoPath[:len(videoPath)-10]
// 	}
// 	videoPath = videoPath + v.Name + ".mp4"
// 	v.OriginalVideoPath = videoPath

// 	// Save the updated `OriginalVideoPath` field to the database
// 	if err = tx.Save(v).Error; err != nil {
// 		return err
// 	}
// 	return
// }
