package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// On-call schedule entity
type OnCallSchedule struct {
	ID       int
	Start    time.Time
	End      time.Time
	Rotation []int // Array of on-call personnel IDs in rotation order
	TimeZone string
}

// On-call personnel entity
type OnCallPersonnel struct {
	ID           int
	Name         string
	Phone        string
	Email        string
	NotifyMethod string // SMS, email, phone call
}

// Incident entity
type Incident struct {
	ID          int
	Severity    string // low, medium, high
	Description string
	Timestamp   time.Time
}

// Notification entity
type Notification struct {
	ID           int
	IncidentID   int
	PersonnelID  int
	Severity     string
	NotifyMethod string
	Timestamp    time.Time
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
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
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

	r.GET("/oncall/personnel", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "personnels",
		})
	})
	r.GET("/incident", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "incidents",
		})
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
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
