package db

import "apartment_rent/models"

func Migrate() error {
	err := dbConn.AutoMigrate(models.User{},
		models.Announcement{}, models.AnnouncementView{},models.Favorite{},models.Review{})
	if err != nil {
		return err
	}
	return nil
}
