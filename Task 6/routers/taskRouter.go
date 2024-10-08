package router

import (
	"task_manager_api/controllers"
	"task_manager_api/middleware"

	"github.com/gin-gonic/gin"
)

func TaskRouter(router *gin.Engine) {
	router.GET("/tasks", middleware.JWTValidation(), middleware.RoleAuth("ADMIN", "USER"), controllers.GetTasks)
	router.GET("/tasks/:id", middleware.JWTValidation(), middleware.RoleAuth("ADMIN", "USER"), controllers.GetTaskById)
	router.POST("/tasks", middleware.JWTValidation(), middleware.RoleAuth("ADMIN"), controllers.AddTask)
	router.PUT("/tasks/:id", middleware.JWTValidation(), middleware.RoleAuth("ADMIN"), controllers.UpdateTask)
	router.PATCH("/tasks/:id", middleware.JWTValidation(), middleware.RoleAuth("ADMIN"), controllers.UpdateSpecificField)
	router.DELETE("/tasks/:id", middleware.JWTValidation(), middleware.RoleAuth("ADMIN"), controllers.DeleteTask)
}
