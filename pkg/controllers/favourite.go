package controllers

import (
	"apartment_rent/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)




// GetFavorites godoc
// @Summary Get favorite announcements
// @Description Retrieves a list of favorite announcements for the authenticated user
// @Tags favorites
// @Accept json
// @Produce json
// @Success 200 {array} models.Announcement "List of favorite announcements"
// @Failure 500 {object} map[string]string "error: failed to get favorites"
// @Router /favorites [get]
// @Security BearerAuth
func GetFavorites(c *gin.Context) {
	userID := c.GetUint(userIDCtx)

	favorites, err := service.GetFavorites(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get favorites",
		})
		return
	}

	c.JSON(http.StatusOK, favorites)
}



// AddToFavorites godoc
// @Summary Get favorite announcements
// @Description Retrieves a list of favorite announcements for the authenticated user
// @Tags favorites
// @Accept json
// @Produce json
// @Success 200 {array} models.Announcement "List of favorite announcements"
// @Failure 500 {object} map[string]string "error: failed to get favorites"
// @Router /favorites [get]
// @Router /announcements/{id}/favorite [post]
// @Security BearerAuth
func AddToFavorites(c *gin.Context) {
	announcementIDStr := c.Param("id")
	
	announcementID, err := strconv.ParseUint(announcementIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid announcement ID",
		})
		return
	}

	userID := c.GetUint(userIDCtx)
	err = service.AddToFavorites(userID, uint(announcementID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to add to favorites",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "added to favorites",
	})
}

