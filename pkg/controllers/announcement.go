package controllers

import (
	"apartment_rent/logger"
	"apartment_rent/models"
	"apartment_rent/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllAnnouncement
// @Summary Get All Announcement
// @Security ApiKeyAuth
// @Tags announcements
// @Description get list of all announcements
// @ID get-all-announcements
// @Produce json
// @Param q query string false "fill if you need search"
// @Success 200 {array} models.Announcement
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /announcements [get]
func GetAllAnnouncement(c *gin.Context) {
	logger.Info.Printf("Client with ip: [%s] requested list of announcements\n", c.ClientIP())
	announcements, err := service.GetAllAnnouncement()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"announcements": announcements,
	})
	logger.Info.Printf("Client with ip: [%s] got list of announcements\n", c.ClientIP())
}

// GetAnnouncementByPrice
// @Summary Get Announcements By Price
// @Security ApiKeyAuth
// @Tags announcements
// @Description get announcements by price
// @ID get-announcement-by-price
// @Produce json
// @Param rooms path integer true "number of price in the announcement"
// @Success 200 {array} models.Announcement
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /announcements/{price} [get]
// Исправленный код контроллера
func GetAnnouncementByPrice(c *gin.Context) {
	prices, err := strconv.Atoi(c.Param("price"))
	if err != nil {
		logger.Error.Printf("[controllers.GetAnnouncementByPrice] invalid path parameter: %s\n", c.Param("price"))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid number of rooms",
		})
		return
	}

	announcements, err := service.GetAnnouncementsByPrice(uint(prices))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"announcements": announcements,
	})
}

// CreateAnnouncement
// @Summary Create Announcement
// @Security ApiKeyAuth
// @Tags announcements
// @Description create new announcement
// @ID create-announcement
// @Accept json
// @Produce json
// @Param input body models.Announcement true "new announcement info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /announcements [post]
func CreateAnnouncement(c *gin.Context) {
	var announcement models.Announcement
	if err := c.BindJSON(&announcement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	userID := c.GetUint(userIDCtx)
	announcement.CreatedBy = int(userID)

	err := service.CreateAnnouncement(announcement)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "annoucement created successfully",
	})

}

// UpdateAnnouncement
// @Summary Update Announcement
// @Security ApiKeyAuth
// @Tags announcements
// @Description Update an existing announcement, only if it was created by the current user
// @ID update-announcement
// @Accept json
// @Produce json
// @Param id path integer true "id of the announcement"
// @Param input body models.AnnouncementFilterForUpdate true "announcement update info"
// @Success 200 {object} models.Announcement
// @Failure 400 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /announcements/{id} [put]
func UpdateAnnouncement(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid announcement id",
		})
		return
	}

	userID := c.GetUint(userIDCtx)
	announcement, err := service.GetAnnouncementByID(id,userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if announcement.CreatedBy != int(userID) {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "you can only update your own announcements",
		})
		return
	}

	var annUpdate models.AnnouncementFilterForUpdate
	if err := c.BindJSON(&annUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input",
		})
		return
	}
	updatedAnnouncement, err := service.UpdateAnnouncement(id, annUpdate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "announcement updated successfully",
		"announcement":  updatedAnnouncement,
	})
}


// DeleteAnnouncement
// @Summary Delete Announcement
// @Security ApiKeyAuth
// @Tags announcements
// @Description delete an existing announcement
// @ID delete-announcement
// @Param id path integer true "id of the announcement to delete"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /announcements/{id} [delete]
func DeleteAnnouncement(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.DeleteAnnouncement] invalid path parameter: %s\n", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid announcement ID",
		})
		return
	}

	err = service.DeleteAnnouncement(uint(id))
	if err != nil {
		logger.Error.Printf("[controllers.DeleteAnnouncement] error deleting announcement: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Announcement deleted successfully",
	})

	logger.Info.Printf("Client with ip: [%s] deleted announcement with ID: %d\n", c.ClientIP(), id)
}

// GetAnnouncementByRooms
// @Summary Get Announcements By Number of Rooms
// @Security ApiKeyAuth
// @Tags announcements
// @Description get announcements by number of rooms
// @ID get-announcement-by-rooms
// @Produce json
// @Param rooms path integer true "number of rooms in the announcement"
// @Success 200 {array} models.Announcement
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /announcements/count_apart/{count_apart} [get]
func GetAnnouncementByRooms(c *gin.Context) {
	rooms, err := strconv.Atoi(c.Param("count_apart"))
	if err != nil {
		logger.Error.Printf("[controllers.GetAnnouncementByRooms] invalid path parameter: %s\n", c.Param("rooms"))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid number of rooms",
		})
		return
	}

	announcements, err := service.GetAnnouncementsByRooms(uint(rooms))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"announcements": announcements,
	})
}

func GetAnnouncementByID(c *gin.Context) {
	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный ID объявления"})
		return
	}

	userID := c.GetUint(userIDCtx)
	announcement, err := service.GetAnnouncementByID(id, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, announcement)
}

