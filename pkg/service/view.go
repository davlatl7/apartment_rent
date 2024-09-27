package service

import (
	"apartment_rent/models"
	
	"apartment_rent/db"

)


func GetViewsReport(announcementID uint) ([]models.AnnouncementView, error) {
	var views []models.AnnouncementView
	err := db.GetDBConn().Where("announcement_id = ?", announcementID).Find(&views).Error
	return views, err
}


