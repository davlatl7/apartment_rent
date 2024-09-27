package service

import (
	"apartment_rent/models"
	"apartment_rent/db"
    "errors"
)

func CreateReview(review models.Review) error {
	return db.GetDBConn().Create(&review).Error
}

func GetReviews(announcementID uint) ([]models.Review, error) {
	var reviews []models.Review

	err := db.GetDBConn().
		Where("announcement_id = ?", announcementID).
		Order("created_at DESC"). 
		Find(&reviews).Error

	if err != nil {
	
		return nil, err
	}

	if len(reviews) == 0 {
		return []models.Review{}, nil
	}

	return reviews, nil
}

func DeleteReview(reviewID uint, userID uint) error {
	var review models.Review

	
	err := db.GetDBConn().First(&review, reviewID).Error
	if err != nil {
		return errors.New("Review not found")
	}

	
	err = db.GetDBConn().Delete(&review).Error
	if err != nil {
		return errors.New("Failed to delete review")
	}

	return nil
}