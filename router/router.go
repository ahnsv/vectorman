package router

import (
	"github.com/ahnsv/vectorman/docs"
	v1 "github.com/ahnsv/vectorman/router/v1"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CreateRouter() *gin.Engine {
	// Create a gin router with default middleware:
	// logger and recovery (crash-free) middleware

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

	return r
}
