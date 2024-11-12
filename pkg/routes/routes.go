package routes

import (
	"golang_api/pkg/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/users", handlers.GetUsers)
	r.GET("/user/:user_id", handlers.GetUser)
	r.POST("/user", handlers.CreateUser)
	return r
}
