package handlers

import (
	"net/http"
	"strconv"

	"golang_api/pkg/models"
	"golang_api/pkg/repositories"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDTO.Usr_password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to hash password"})
		return
	}

	userRepo := repositories.NewUserRepository()
	user := models.Tb_User{
		Usr_name:     userDTO.Usr_name,
		Usr_email:    userDTO.Usr_email,
		Usr_password: string(hashedPassword),
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

// UpdateUser godoc
// @Summary Update an existing user
// @Description UpdateUser updates an existing user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param user_id path int true "User ID"
// @Param user body models.UserDTO true "User DTO"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /user/{user_id} [put]
func UpdateUser(c *gin.Context) {
	userIDParam := c.Param("user_id")
	userRepo := repositories.NewUserRepository()

	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid user ID"})
		return
	}

	var userDTO models.UserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid input"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDTO.Usr_password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to hash password"})
		return
	}

	user := models.Tb_User{
		Usr_id:       userID,
		Usr_name:     userDTO.Usr_name,
		Usr_email:    userDTO.Usr_email,
		Usr_password: string(hashedPassword),
	}

	err = userRepo.Update(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	userResponse := models.UserResponse{
		Usr_id:    user.Usr_id,
		Usr_name:  user.Usr_name,
		Usr_email: user.Usr_email,
	}

	c.JSON(http.StatusOK, userResponse)
}

// DeleteUser godoc
// @Summary Delete a user by ID
// @Description DeleteUser deletes a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param user_id path int true "User ID"
// @Success 204
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /user/{user_id} [delete]
func DeleteUser(c *gin.Context) {
	userIDParam := c.Param("user_id")
	userRepo := repositories.NewUserRepository()

	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid user ID"})
		return
	}

	err = userRepo.Delete(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	c.Status(http.StatusNoContent)
}
