package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang_api/pkg/models"
)

// GetUsers godoc
// @Summary Get list of users aaa
// @Description GetUsers returns a list of users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Failure 500 {object} string
// @Router /users [get]
func GetUsers(c *gin.Context) {
	users, err := fetchUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func fetchUsers() ([]string, error) {
	return []string{"Alice", "Bob", "Charlie"}, nil
}
