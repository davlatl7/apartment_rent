package repository

import (
	"apartment_rent/db"
	"apartment_rent/logger"
	"apartment_rent/models"
	"errors"
	"gorm.io/gorm"
)

func GetAllAnnouncement() (announcements []models.Announcement, err error) {
	err = db.GetDBConn().Find(&announcements).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllAnnouncement] error getting all announcements: %s\n", err.Error())
		return nil, translateError(err)
	}

	for i := range announcements {
		var viewCount int64
		announcementID := announcements[i].ID

		err = db.GetDBConn().Model(&models.AnnouncementView{}).Where("announcement_id = ?", announcementID).Count(&viewCount).Error
		if err != nil {
			logger.Error.Printf("[repository.GetAllAnnouncement] error counting views for announcement ID %d: %v\n", announcementID, err)
			return nil, translateError(err)
		}

		announcements[i].ViewCount = viewCount
	}

	return announcements, nil
}

func GetAnnouncementsByPrice(rooms uint) (announcements []models.Announcement, err error) {
    err = db.GetDBConn().Where("price = ?", rooms).Find(&announcements).Error
    if err != nil {
        logger.Error.Printf("[repository.GetAnnouncementsByRooms] error getting announcements by count_apart: %v\n", err)
        return nil, translateError(err)
    }

    return announcements, nil
}

func CreateAnnouncement(announcement models.Announcement) (err error) {
	if err = db.GetDBConn().Create(&announcement).Error; err != nil {
		logger.Error.Printf("[repository.CreateAnnouncement] error creating announcement: %v\n", err)
		return translateError(err)
	}

	return nil
}


func DeleteAnnouncement(id uint) error {
	err := db.GetDBConn().Where("id = ?", id).Delete(&models.Announcement{}).Error
	if err != nil {
		logger.Error.Printf("[repository.DeleteAnnouncement] error deleting announcement with ID %d: %v\n", id, err)
		return translateError(err)
	}

	return nil
}


func GetAnnouncementsByRooms(rooms uint) (announcements []models.Announcement, err error) {
    err = db.GetDBConn().Where("count_apart = ?", rooms).Find(&announcements).Error
    if err != nil {
        logger.Error.Printf("[repository.GetAnnouncementsByRooms] error getting announcements by count_apart: %v\n", err)
        return nil, translateError(err)
    }

    return announcements, nil
}

func GetAnnouncementByID(id int, userID uint) (*models.Announcement, error) {
	var announcement models.Announcement

	// Получаем объявление по ID
	err := db.GetDBConn().Where("id = ?", id).First(&announcement).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("объявление не найдено")
		}
		return nil, err
	}

	// Проверяем, смотрел ли пользователь это объявление
	var view models.AnnouncementView
	err = db.GetDBConn().Where("announcement_id = ? AND user_id = ?", id, userID).First(&view).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Пользователь еще не просматривал объявление, увеличиваем счётчик
		announcement.ViewCount++

		// Сохраняем обновлённый счётчик просмотров
		err = db.GetDBConn().Model(&announcement).Update("view_count", announcement.ViewCount).Error
		if err != nil {
			return nil, err
		}

		// Создаём запись о просмотре для данного пользователя
		view = models.AnnouncementView{
			AnnouncementID: uint(id),
			UserID:         userID,
		}
		err = db.GetDBConn().Create(&view).Error
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		// В случае другой ошибки
		return nil, err
	}

	// Возвращаем объявление с актуальным количеством просмотров
	return &announcement, nil
}




func UpdateAnnouncement(id int, annUpdate models.AnnouncementFilterForUpdate) (*models.Announcement, error) {
	updates := map[string]interface{}{}

	if annUpdate.Price != 0 {
		updates["price"] = annUpdate.Price
	}
	if annUpdate.CountApart != 0 {
		updates["count_apart"] = annUpdate.CountApart
	}
	if annUpdate.Floor != 0 {
		updates["floor"] = annUpdate.Floor
	}
	if annUpdate.SquareMeters != 0 {
		updates["square_meters"] = annUpdate.SquareMeters
	}
	if annUpdate.District != "" {
		updates["district"] = annUpdate.District
	}
	if annUpdate.Pets != "" {
		updates["pets"] = annUpdate.Pets
	}
	if annUpdate.TypeOfDevelopment != "" {
		updates["typeofdevelopment"] = annUpdate.TypeOfDevelopment
	}
	if annUpdate.Comment != "" {
		updates["comment"] = annUpdate.Comment
	}
	if annUpdate.Phone != "" {
		updates["phone"] = annUpdate.Phone
	}

	if len(updates) > 0 {
		err := db.GetDBConn().Model(&models.Announcement{}).Where("id = ?", id).Updates(updates).Error
		if err != nil {
			return nil, err
		}
	}

	var announcement models.Announcement
	err := db.GetDBConn().Where("id = ?", id).First(&announcement).Error
	if err != nil {
		return nil, err
	}

	return &announcement, nil
}
