package routes

import (
	"golang_api/pkg/handlers"
	"golang_api/pkg/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.CORSMiddleware())

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//user
	r.GET("/users", handlers.GetUsers)
	r.GET("/user/:user_id", handlers.GetUser)
	r.POST("/user", handlers.CreateUser)
	r.PUT("/user/:user_id", handlers.UpdateUser)
	r.DELETE("/user/:user_id", handlers.DeleteUser)

	//task
	r.GET("/tasks", handlers.GetTasks)
	r.GET("/task/:task_id", handlers.GetTask)
	r.POST("/task", handlers.CreateTask)
	r.PUT("/task/:task_id", handlers.UpdateTask)
	r.DELETE("/task/:task_id", handlers.DeleteTask)

	//status
	r.GET("/status", handlers.GetAllStatus)
	r.GET("/status/:tskst_id", handlers.GetStatus)

	//priorities
	r.GET("/priorities", handlers.GetPriorities)
	r.GET("/priority/:tskpr_id", handlers.GetPriority)

	r.POST("/login", handlers.Login)
	return r
}
