package handlers

import (
    "github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
    users := []string{"Alice", "Bob", "Charlie"}
    c.JSON(200, users)
}
