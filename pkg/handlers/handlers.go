package handlers

import (
	"net/http"
	"strconv"

	"golang_api/pkg/models"
	"golang_api/pkg/repositories"

	"github.com/gin-gonic/gin"
)

// GetUser godoc
// @Summary Get user by ID
// @Description GetUser returns a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param user_id path int true "User ID"
// @Success 200 {object} repositories.User
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /user/{user_id} [get]
func GetUser(c *gin.Context) {
	userIDParam := c.Param("user_id")
	userRepo := repositories.NewUserRepository()

	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid user ID"})
		return
	}

	user, err := userRepo.Select(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetUsers godoc
// @Summary Get all users
// @Description GetUsers returns all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} repositories.User
// @Failure 500 {object} models.ErrorResponse
// @Router /users [get]
func GetUsers(c *gin.Context) {
	userRepo := repositories.NewUserRepository()
	users, err := userRepo.Select()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// CreateUser godoc
// @Summary Create a new user
// @Description CreateUser creates a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.UserDTO true "User DTO"
// @Success 201 {object} models.UserResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /user [post]
func CreateUser(c *gin.Context) {
	var userDTO models.UserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid input"})
		return
	}

	userRepo := repositories.NewUserRepository()
	user := models.Tb_User{
		Usr_name:     userDTO.Usr_name,
		Usr_email:    userDTO.Usr_email,
		Usr_password: userDTO.Usr_password,
	}

	userID, err := userRepo.Insert(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	userResponse := models.UserResponse{
		Usr_id:    userID,
		Usr_name:  user.Usr_name,
		Usr_email: user.Usr_email,
	}

	c.JSON(http.StatusCreated, userResponse)
}
