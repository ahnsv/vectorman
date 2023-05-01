package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ahnsv/vectorman/pkg/app"
	"github.com/ahnsv/vectorman/pkg/entities"
	"github.com/gin-gonic/gin"
)

var incidents []entities.Incident

// @Summary Get all incidents
// @Produce json
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /api/v1/incidents [get]
func GetIncidents(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, 200, incidents)
}

// @Summary Get incident by ID
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /api/v1/incidents/{id} [get]
func GetIncidentByID(c *gin.Context) {
	appG := app.Gin{C: c}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	for i := range incidents {
		if incidents[i].ID == id {
			appG.Response(http.StatusOK, 200, incidents[i])
			return
		}
	}
	appG.Response(http.StatusNotFound, 404, gin.H{"error": "Incident not found"})
}

// @Summary Create a new incident
// @Produce json
// @Param request body entities.Incident true "Incident"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /api/v1/incidents [post]
func CreateIncident(c *gin.Context) {
	appG := app.Gin{C: c}
	var newIncident entities.Incident
	err := c.BindJSON(&newIncident)
	if err != nil {
		println(err.Error())
		appG.Response(http.StatusBadRequest, 400, nil)
		return
	}
	incidents = append(incidents, newIncident)

	// create notification for new incident
	var newNotification entities.Notification
	newNotification.ID = len(notifications) + 1
	newNotification.IncidentID = newIncident.ID
	newNotification.PersonnelID = 1 // entities.GetCurrentOncallSchedule().Rotation[0]
	newNotification.Severity = newIncident.Severity
	newNotification.Timestamp = time.Now()

	newNotification.Send()

	appG.Response(http.StatusOK, 200, gin.H{"status": "OK"})
}

// @Summary Update an incident
// @Produce json
// @Param id path int true "ID"
// @Param request body entities.Incident true "Incident"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /api/v1/incidents/{id} [put]
func UpdateIncident(c *gin.Context) {
	appG := app.Gin{C: c}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var updatedIncident entities.Incident
	err = c.BindJSON(&updatedIncident)
	if err != nil {
		appG.Response(http.StatusBadRequest, 400, nil)
		return
	}
	for i := range incidents {
		if incidents[i].ID == id {
			incidents[i] = updatedIncident
			appG.Response(http.StatusOK, 200, gin.H{"status": "OK"})
			return
		}
	}
	appG.Response(http.StatusNotFound, 404, gin.H{"error": "Incident not found"})
}

// @Summary Delete an incident
// @Produce json
// @Param id path int true "ID"
// @Success 204 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /api/v1/incidents/{id} [delete]
func DeleteIncident(c *gin.Context) {
	appG := app.Gin{C: c}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	for i := range incidents {
		if incidents[i].ID == id {
			incidents = append(incidents[:i], incidents[i+1:]...)
			appG.Response(http.StatusAccepted, 204, gin.H{"status": "OK"})
			return
		}
	}
	appG.Response(http.StatusNotFound, 404, gin.H{"error": "Incident not found"})
}
