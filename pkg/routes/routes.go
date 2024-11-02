package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
    "golang_api/pkg/handlers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    // Swagger endpoint
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    r.GET("/users", handlers.GetUsers)
    return r
}
