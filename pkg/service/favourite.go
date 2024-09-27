package service

import (
	"apartment_rent/models"
	"apartment_rent/db"
)



func AddToFavorites(userID uint, announcementID uint) error {
	favorite := models.Favorite{
		UserID:        userID,
		AnnouncementID: announcementID,
	}

	return db.GetDBConn().Create(&favorite).Error
}

func GetFavorites(userID uint) ([]models.Announcement, error) {
	var favorites []models.Announcement
	err := db.GetDBConn().Table("favorites").
		Select("announcements.*").
		Joins("join announcements on announcements.id = favorites.announcement_id").
		Where("favorites.user_id = ?", userID).
		Scan(&favorites).Error
	return favorites, err
}