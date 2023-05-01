package v1

import (
	"net/http"
	"strconv"

	"github.com/ahnsv/vectorman/pkg/app"
	"github.com/ahnsv/vectorman/pkg/entities"
	"github.com/gin-gonic/gin"
)

var notifications []entities.Notification

// @Summary Get all notifications
// @Produce json
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /api/v1/notifications [get]
func GetNotifications(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, 200, notifications)
}

// @Summary Get notification by ID
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /api/v1/notifications/{id} [get]
func GetNotificationByID(c *gin.Context) {
	appG := app.Gin{C: c}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appG.Response(http.StatusBadRequest, 400, gin.H{"error": "Invalid ID"})
		return
	}
	for i := range notifications {
		if notifications[i].ID == id {
			appG.Response(http.StatusOK, 200, notifications[i])
			return
		}
	}
	appG.Response(http.StatusNotFound, 404, gin.H{"error": "Notification not found"})
}

// @Summary Create a new notification
// @Produce json
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /api/v1/notifications [post]
func CreateNotification(c *gin.Context) {
	appG := app.Gin{C: c}
	var newNotification entities.Notification
	err := c.BindJSON(&newNotification)
	if err != nil {
		appG.Response(http.StatusBadRequest, 400, nil)
		return
	}
	notifications = append(notifications, newNotification)
}

// @Summary Update a notification
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /api/v1/notifications/{id} [put]
func UpdateNotification(c *gin.Context) {
	appG := app.Gin{C: c}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appG.Response(http.StatusBadRequest, 400, gin.H{"error": "Invalid ID"})
		return
	}
	var updatedNotification entities.Notification
	err = c.BindJSON(&updatedNotification)
	if err != nil {
		appG.Response(http.StatusBadRequest, 400, nil)
		return
	}
	for i := range notifications {
		if notifications[i].ID == id {
			notifications[i] = updatedNotification
			appG.Response(http.StatusOK, 200, notifications[i])
			return
		}
	}
	appG.Response(http.StatusNotFound, 404, gin.H{"error": "Notification not found"})
}

// @Summary Delete a notification
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /api/v1/notifications/{id} [delete]
func DeleteNotification(c *gin.Context) {
	appG := app.Gin{C: c}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appG.Response(http.StatusBadRequest, 400, gin.H{"error": "Invalid ID"})
		return
	}
	for i := range notifications {
		if notifications[i].ID == id {
			notifications = append(notifications[:i], notifications[i+1:]...)
			appG.Response(http.StatusOK, 200, nil)
			return
		}
	}
	appG.Response(http.StatusNotFound, 404, gin.H{"error": "Notification not found"})
}
