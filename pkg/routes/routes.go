package routes

import (
    "github.com/gin-gonic/gin"
    "golang_api/pkg/handlers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.GET("/users", handlers.GetUsers)
    return r
}
