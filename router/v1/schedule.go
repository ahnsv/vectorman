package v1

import (
	"net/http"
	"strconv"

	"github.com/ahnsv/vectorman/pkg/app"
	"github.com/ahnsv/vectorman/pkg/e"
	"github.com/ahnsv/vectorman/pkg/entities"
	"github.com/gin-gonic/gin"
)

var onCallSchedules []entities.OnCallSchedule

// @Summary: Define endpoint for creating a new On-call schedule
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /api/v1/oncall/schedule [post]
func CreateOncallSchedule(c *gin.Context) {
	appG := app.Gin{C: c}

	var newSchedule entities.OnCallSchedule
	err := c.BindJSON(&newSchedule)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	onCallSchedules = append(onCallSchedules, newSchedule)
	appG.Response(http.StatusOK, 200, gin.H{"status": "OK"})
}

// @Summary: Define endpoint for updating an existing On-call schedule
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /api/v1/oncall/schedule/{id} [put]
func UpdateOncallSchedule(c *gin.Context) {
	appG := app.Gin{C: c}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, gin.H{"error": "Invalid ID"})
		return
	}
	var updatedSchedule entities.OnCallSchedule
	err = c.BindJSON(&updatedSchedule)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, gin.H{"error": "Invalid JSON"})
		return
	}
	for i := range onCallSchedules {
		if onCallSchedules[i].ID == id {
			onCallSchedules[i] = updatedSchedule
			appG.Response(http.StatusOK, e.SUCCESS, gin.H{"status": "OK"})
			return
		}
	}
	appG.Response(http.StatusNotFound, e.ERROR_NOT_EXIST_SCHDULE, gin.H{"error": "Schedule not found"})
}

// @Summary: Define endpoint for getting a list of all On-call schedules
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /api/v1/oncall/schedule [get]
func GetOncallSchedule(c *gin.Context) {
	c.JSON(http.StatusOK, onCallSchedules)
}

// @Summary: Define endpoint for getting a specific On-call schedule by ID
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /api/v1/oncall/schedule/{id} [get]
func GetOncallScheduleByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	for _, schedule := range onCallSchedules {
		if schedule.ID == id {
			c.JSON(http.StatusOK, schedule)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
}

// @Summary: Define endpoint for deleting an existing On-call schedule
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /api/v1/oncall/schedule/{id} [delete]
func DeleteOncallSchedule(c *gin.Context) {
	appG := app.Gin{C: c}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, gin.H{"error": "Invalid ID"})
		return
	}
	for i := range onCallSchedules {
		if onCallSchedules[i].ID == id {
			onCallSchedules = append(onCallSchedules[:i], onCallSchedules[i+1:]...)
			appG.Response(http.StatusOK, e.SUCCESS, gin.H{"status": "OK"})
			return
		}
	}
	appG.Response(http.StatusNotFound, e.ERROR_NOT_EXIST_SCHDULE, gin.H{"error": "Schedule not found"})
}
