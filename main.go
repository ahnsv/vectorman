package main

import (
	"time"

	docs "github.com/ahnsv/vectorman/docs"
	"github.com/ahnsv/vectorman/pkg/entities"
	v1 "github.com/ahnsv/vectorman/router/v1"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

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
		oncall := apiv1.Group("/oncall")
		{
			oncall.GET("/schedule", v1.GetOncallSchedule)
			oncall.GET("/schedule/:id", v1.GetOncallScheduleByID)
			oncall.POST("/schedule", v1.CreateOncallSchedule)
			oncall.PUT("/schedule/:id", v1.UpdateOncallSchedule)
			oncall.DELETE("/schedule/:id", v1.DeleteOncallSchedule)
			oncall.GET("/personnel", v1.GetOncallPersonnel)
			oncall.GET("/personnel/:id", v1.GetOncallPersonnelByID)
			oncall.POST("/personnel", v1.CreateOncallPersonnel)
			oncall.PUT("/personnel/:id", v1.UpdateOncallPersonnel)
			oncall.DELETE("/personnel/:id", v1.DeleteOncallPersonnel)
		}

		incident := apiv1.Group("/incidents")
		{
			incident.GET("/", v1.GetIncidents)
			incident.GET("/:id", v1.GetIncidentByID)
			incident.POST("/", v1.CreateIncident)
			incident.PUT("/:id", v1.UpdateIncident)
			incident.DELETE("/:id", v1.DeleteIncident)
		}

		notification := apiv1.Group("/notifications")
		{
			notification.GET("/", v1.GetNotifications)
			notification.GET("/:id", v1.GetNotificationByID)
			notification.POST("/", v1.CreateNotification)
			notification.PUT("/:id", v1.UpdateNotification)
			notification.DELETE("/:id", v1.DeleteNotification)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
