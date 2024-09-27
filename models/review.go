package models

 import "time"

 type Review struct {
	ID            uint      `json:"id" gorm:"primary_key"`
	AnnouncementID uint     `json:"announcement_id"`
	UserID        uint      `json:"user_id"`
	Content       string    `json:"content"`
	Rating        int       `json:"rating"`
	CreatedAt     time.Time `json:"created_at"`
}


