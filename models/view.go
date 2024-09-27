package models 
import "time" 

type View struct {
    ID            uint   `json:"id"`
    AnnouncementID uint   `json:"announcement_id"`
    UserID        uint   `json:"user_id"`
    CreatedAt     time.Time `json:"created_at"`
}