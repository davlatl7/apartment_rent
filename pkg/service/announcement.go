package service

import (
	"apartment_rent/models"
	"apartment_rent/pkg/repository"
	
	
)

func GetAllAnnouncement() (announcements []models.Announcement, err error) {
	announcements, err = repository.GetAllAnnouncement()
	if err != nil {
		return nil, err
	}

	return announcements, nil
}

func GetAnnouncementsByPrice(price uint) (announcements []models.Announcement, err error) {
	announcements, err = repository.GetAnnouncementsByPrice(price)
	if err != nil {
		return nil, err
	}

	return announcements, nil
}
func CreateAnnouncement(announcement models.Announcement) error {

	err := repository.CreateAnnouncement(announcement)
	if err != nil {
		return err
	}

	return nil
}

func DeleteAnnouncement(id uint) error {

	err := repository.DeleteAnnouncement(id)
	if err != nil {
		return err
	}

	return nil
}


func GetAnnouncementsByRooms(rooms uint) (announcements []models.Announcement, err error) {
	announcements, err = repository.GetAnnouncementsByRooms(rooms)
	if err != nil {
		return nil, err
	}

	return announcements, nil
}

func GetAnnouncementByID(id int, userID uint) (*models.Announcement, error) {
	announcement, err := repository.GetAnnouncementByID(id, userID)
	if err != nil {
		return nil, err
	}

	return announcement, nil
}

func UpdateAnnouncement(id int, annUpdate models.AnnouncementFilterForUpdate) (*models.Announcement, error) {
	announcement, err := repository.UpdateAnnouncement(id, annUpdate)
	if err != nil {
		return nil, err
	}
	return announcement, nil
}

