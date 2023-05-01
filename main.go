package main

import (
	"net/http"
	"time"

	docs "github.com/ahnsv/vectorman/docs"
	"github.com/ahnsv/vectorman/pkg/entities"
	v1 "github.com/ahnsv/vectorman/router/v1"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

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
	schedule  *entities.OnCallSchedule
	personnel []*entities.OnCallPersonnel
}

// Incident aggregate root
type IncidentRoot struct {
	incident *Incident
}

// Notification aggregate root
type NotificationRoot struct {
	notification *Notification
}

var notifications []Notification

func main() {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Title = "Vectorman API"
	docs.SwaggerInfo.Description = "Vectorman API"

	apiv1 := r.Group("/api/v1")
	{
		eg := apiv1.Group("/oncall")
		{
			eg.GET("/schedule", v1.GetOncallSchedule)
			eg.GET("/schedule/:id", v1.GetOncallScheduleByID)
			eg.POST("/schedule", v1.CreateOncallSchedule)
			eg.PUT("/schedule/:id", v1.UpdateOncallSchedule)
			eg.DELETE("/schedule/:id", v1.DeleteOncallSchedule)
			eg.GET("/personnel", v1.GetOncallPersonnel)
			eg.GET("/personnel/:id", v1.GetOncallPersonnelByID)
			eg.POST("/personnel", v1.CreateOncallPersonnel)
			eg.PUT("/personnel/:id", v1.UpdateOncallPersonnel)
			eg.DELETE("/personnel/:id", v1.DeleteOncallPersonnel)
		}

		incident := apiv1.Group("/incidents")
		{
			incident.GET("/", v1.GetIncidents)
			incident.GET("/:id", v1.GetIncidentByID)
			incident.POST("/", v1.CreateIncident)
			incident.PUT("/:id", v1.UpdateIncident)
			incident.DELETE("/:id", v1.DeleteIncident)
		}
	}

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
