package models

type Favorite struct {
	ID            uint `json:"id" gorm:"primary_key"`
	UserID        uint `json:"user_id"`
	AnnouncementID uint `json:"announcement_id"`
}