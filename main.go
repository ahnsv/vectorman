package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// On-call schedule entity
type OnCallSchedule struct {
	ID       int       `json:"id"`
	Start    time.Time `json:"start" binding:"required"`
	End      time.Time `json:"end" binding:"required"`
	Rotation []int     `json:"rotation" binding:"required"`
	TimeZone string    `json:"timezone" binding:"required"`
}

// On-call personnel entity
type OnCallPersonnel struct {
	ID           int    `json:"id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Phone        string `json:"phone" binding:"required"`
	Email        string `json:"email" binding:"required"`
	NotifyMethod string `json:"notify_method" binding:"required"`
}

// Incident entity
type Incident struct {
	ID          int       `json:"id" binding:"required"`
	Severity    string    `json:"severity" binding:"required"` // low, medium, high
	Description string    `json:"description" binding:"required"`
	Status      string    `json:"status" binding:"required"` // open, closed
	Timestamp   time.Time `json:"timestamp" binding:"required"`
}

// Notification entity
type Notification struct {
	ID           int       `json:"id" binding:"required"`
	IncidentID   int       `json:"incident_id" binding:"required"`
	PersonnelID  int       `json:"personnel_id" binding:"required"`
	Severity     string    `json:"severity" binding:"required"`
	NotifyMethod string    `json:"notify_method" binding:"required"`
	Timestamp    time.Time `json:"timestamp" binding:"required"`
}

// On-call schedule aggregate root
type OnCallScheduleRoot struct {
	schedule  *OnCallSchedule
	personnel []*OnCallPersonnel
}

// Incident aggregate root
type IncidentRoot struct {
	incident *Incident
}

// Notification aggregate root
type NotificationRoot struct {
	notification *Notification
}

var onCallSchedules []OnCallSchedule
var onCallPersonnel []OnCallPersonnel
var incidents []Incident
var notifications []Notification

func main() {
	r := gin.Default()

	// Define endpoint for creating a new On-call schedule
	r.POST("/oncall/schedule", func(c *gin.Context) {
		var newSchedule OnCallSchedule
		err := c.BindJSON(&newSchedule)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}
		onCallSchedules = append(onCallSchedules, newSchedule)
		c.JSON(http.StatusCreated, gin.H{"status": "OK"})
	})

	// Define endpoint for updating an existing On-call schedule
	r.PUT("/oncall/schedule/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		var updatedSchedule OnCallSchedule
		err = c.BindJSON(&updatedSchedule)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}
		for i := range onCallSchedules {
			if onCallSchedules[i].ID == id {
				onCallSchedules[i] = updatedSchedule
				c.JSON(http.StatusOK, gin.H{"status": "OK"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
	})

	// Define endpoint for getting a list of all On-call schedules
	r.GET("/oncall/schedule", func(c *gin.Context) {
		c.JSON(http.StatusOK, onCallSchedules)
	})

	// Define endpoint for getting a specific On-call schedule by ID
	r.GET("/oncall/schedule/:id", func(c *gin.Context) {
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
	})

	r.POST("/oncall/personnel", func(c *gin.Context) {
		var newPersonnel OnCallPersonnel
		err := c.BindJSON(&newPersonnel)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}
		onCallPersonnel = append(onCallPersonnel, newPersonnel)
		c.JSON(http.StatusCreated, gin.H{"status": "OK"})
	})
	r.GET("/oncall/personnel", func(c *gin.Context) {
		c.JSON(http.StatusOK, onCallPersonnel)
	})
	r.PUT("/oncall/personnel/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		var updatedPersonnel OnCallPersonnel
		err = c.BindJSON(&updatedPersonnel)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}
		for i := range onCallPersonnel {
			if onCallPersonnel[i].ID == id {
				onCallPersonnel[i] = updatedPersonnel
				c.JSON(http.StatusOK, gin.H{"status": "OK"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Personnel not found"})
	})
	r.GET("/oncall/personnel/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		for _, personnel := range onCallPersonnel {
			if personnel.ID == id {
				c.JSON(http.StatusOK, personnel)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Personnel not found"})
	})
	r.DELETE("/oncall/personnel/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		for i, personnel := range onCallPersonnel {
			if personnel.ID == id {
				onCallPersonnel = append(onCallPersonnel[:i], onCallPersonnel[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"status": "OK"})
				return
			}
		}
	})

	// Define endpoint for creating a new incident
	r.POST("/incidents", func(c *gin.Context) {
		var newIncident Incident
		if err := c.ShouldBindJSON(&newIncident); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newIncident.ID = len(incidents) + 1
		newIncident.Status = "Open"
		incidents = append(incidents, newIncident)
		c.JSON(http.StatusCreated, newIncident)
	})

	// Define endpoint for retrieving a specific incident by ID
	r.GET("/incidents/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, incident := range incidents {
			if strconv.Itoa(incident.ID) == id {
				c.JSON(http.StatusOK, incident)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Incident not found"})
	})

	// Define endpoint for updating an existing incident
	r.PUT("/incidents/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, incident := range incidents {
			if strconv.Itoa(incident.ID) == id {
				var updatedIncident Incident
				if err := c.ShouldBindJSON(&updatedIncident); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}
				updatedIncident.ID = incident.ID
				incidents[i] = updatedIncident
				c.JSON(http.StatusOK, updatedIncident)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Incident not found"})
	})

	// Define endpoint for deleting an existing incident
	r.DELETE("/incidents/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, incident := range incidents {
			if strconv.Itoa(incident.ID) == id {
				incidents = append(incidents[:i], incidents[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"status": "OK"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Incident not found"})
	})

	r.GET("/notification", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "notifications",
		})
	})

	r.GET("/notification/personnel", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "notification personnels",
		})
	})
	r.GET("/notification/incident", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "incidents",
		})
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
