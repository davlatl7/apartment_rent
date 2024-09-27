package controllers

import (
	"apartment_rent/pkg/service"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)





// GetViewsReport godoc
// @Summary Get views report for an announcement
// @Description Retrieves the total number of views for a specific announcement
// @Tags view annoncements
// @Accept json
// @Produce json
// @Param id path uint true "Announcement ID"
// @Success 200 {object} []models.View "List of views"
// @Failure 400 {object} map[string]string "error: invalid announcement ID"
// @Failure 500 {object} map[string]string "error: failed to get report"
// @Router /announcements/{id}/views/report [get]
// @Security BearerAuth
func GetViewsReport(c *gin.Context) {
	// Получаем строковый ID объявления из параметров
	announcementIDStr := c.Param("id")

	// Преобразуем строковый ID в uint
	announcementID, err := strconv.ParseUint(announcementIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid announcement ID",
		})
		return
	}

	// Получаем отчет о просмотрах
	views, err := service.GetViewsReport(uint(announcementID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get report",
		})
		return
	}

	// Возвращаем отчет
	c.JSON(http.StatusOK, views)
}

