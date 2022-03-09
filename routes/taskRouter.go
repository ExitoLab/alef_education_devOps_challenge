package routes

import (

	controller "alef_education_devops_challenge/controllers"

	"github.com/gin-gonic/gin"
)

//TaskRouter function
func TaskRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/api/v1/tasks", controller.AddTask())
	incomingRoutes.GET("/api/v1/tasks/:task_id", controller.GetTaskByTaskID())
	incomingRoutes.GET("/api/v1/tasks", controller.GetAllTasks())
	incomingRoutes.DELETE("/api/v1/tasks/:task_id", controller.DeleteTaskByTaskID())
	incomingRoutes.PATCH("/api/v1/tasks/:task_id", controller.UpdateTaskByTaskID())
	incomingRoutes.GET("/healthz", controller.HealthCheck())
}