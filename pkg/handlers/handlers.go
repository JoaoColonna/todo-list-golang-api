package handlers

import (
    "github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Summary Get list of users
// @Description GetUsers returns a list of users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Router /users [get]
func GetUsers(c *gin.Context) {
    users := []string{"Alice", "Bob", "Charlie"}
    c.JSON(200, users)
}