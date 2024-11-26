package routes

import (
	"golang_api/pkg/handlers"
	middlewares "golang_api/pkg/middlewares"

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

	//category
	r.GET("/categories", handlers.GetCategorys)
	r.GET("/category/:cat_id", handlers.GetCategory)
	r.POST("/category", handlers.CreateCategory)
	r.PUT("/category/:cat_id", handlers.UpdateCategory)
	r.DELETE("/category/:cat_id", handlers.DeleteCategory)

	//TaskCategory
	r.GET("/taskcategory", handlers.GetTaskCategories)
	r.GET("/taskcategory/:tsk_id", handlers.GetTaskCategory)
	r.POST("/taskcategory", handlers.CreateTaskCategory)
	r.PUT("/taskcategory/:tsk_id", handlers.UpdateTaskCategory)
	r.DELETE("/taskcategory/:tsk_id", handlers.DeleteTaskCategory)

	r.POST("/login", handlers.Login)
	return r
}
