package controllers

import (
	"apartment_rent/models"
	"apartment_rent/logger"
	"apartment_rent/pkg/service"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

// CreateReview godoc
// @Summary Create a review for an announcement
// @Description Adds a review to a specific announcement
// @Tags review
// @Accept json
// @Produce json
// @Param id path uint true "Announcement ID"
// @Param review body models.Review true "Review Body"
// @Success 201 {object} map[string]string "message: review created successfully"
// @Failure 400 {object} map[string]string "error: invalid request"
// @Failure 500 {object} map[string]string "error: failed to create review"
// @Router /announcements/{id}/review [post]
// @Security BearerAuth
func CreateReview(c *gin.Context) {
	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := review.UserID
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	err := service.CreateReview(review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create review"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Review created successfully"})
}

// GetReviews godoc
// @Summary Create a review for an announcement
// @Description Adds a review to a specific announcement
// @Tags review
// @Accept json
// @Produce json
// @Param id path uint true "Announcement ID"
// @Param review body models.Review true "Review Body"
// @Success 201 {object} map[string]string "message: review created successfully"
// @Failure 400 {object} map[string]string "error: invalid request"
// @Failure 500 {object} map[string]string "error: failed to create review"
// @Router /announcements/reviews/{id} [get]
func GetReviews(c *gin.Context) {
	announcementIDStr := c.Param("id")

	announcementID, err := strconv.ParseUint(announcementIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid announcement ID"})
		return
	}
	reviews, err := service.GetReviews(uint(announcementID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get reviews"})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

func DeleteReview(c *gin.Context) {
	
	reviewIDStr := c.Param("id")
	reviewID, err := strconv.ParseUint(reviewIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
		return
	}
	userID := c.GetUint("user_id")

	err = service.DeleteReview(uint(reviewID), userID)
	if err != nil {
		logger.Error.Printf("Failed to delete review with ID: %d, error: %v", reviewID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	c.JSON(http.StatusOK, gin.H{"message": "Review deleted successfully"})
}
