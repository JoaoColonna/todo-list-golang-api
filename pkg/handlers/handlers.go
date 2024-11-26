package handlers

import (
	"net/http"
	"strconv"

	"golang_api/pkg/models"
	"golang_api/pkg/repositories"
	"golang_api/pkg/utils"

	"time"

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

// Login godoc
// @Summary Authenticate a user
// @Description Login authenticates a user
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body models.UserDTO true "User DTO"
// @Success 200 {object} models.LoginResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /login [post]
func Login(c *gin.Context) {
	var user models.UserDTO

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid data"})
		return
	}

	userRepository := repositories.NewUserRepository()
	dbUser, err := userRepository.SelectByEmail(user.Usr_email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Invalid credentials"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Usr_password), []byte(user.Usr_password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(dbUser.Usr_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Error generating token"})
		return
	}

	userResponse := models.UserResponse{
		Usr_id:    dbUser.Usr_id,
		Usr_name:  dbUser.Usr_name,
		Usr_email: dbUser.Usr_email,
	}

	loginResponse := models.LoginResponse{
		Token:        token,
		UserResponse: userResponse,
	}

	c.JSON(http.StatusOK, loginResponse)
}

// GetTask godoc
// @Summary Get task by ID
// @Description GetTask returns a task by ID
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param task_id path int true "Task ID"
// @Success 200 {object} models.Tb_Task
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /task/{task_id} [get]
func GetTask(c *gin.Context) {
	taskIDParam := c.Param("task_id")
	taskRepo := repositories.NewTaskRepository()

	taskID, err := strconv.Atoi(taskIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	task, err := taskRepo.Select(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// GetTasks godoc
// @Summary Get all task
// @Description GetTask returns all task
// @Tags tasks
// @Accept  json
// @Produce  json
// @Success 200 {array} repositories.TaskRepository
// @Failure 500 {object} models.ErrorResponse
// @Router /tasks [get]
func GetTasks(c *gin.Context) {
	taskRepo := repositories.NewTaskRepository()
	tasks, err := taskRepo.Select()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// CreateTask godoc
// @Summary Create a new task
// @Description CreateTask creates a new task
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param task body models.Task_Request true "Task_Request"
// @Success 201 {object} models.Tb_Task
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /task [post]
func CreateTask(c *gin.Context) {
	var tb_Task models.Task_Request
	if err := c.ShouldBindJSON(&tb_Task); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid input"})
		return
	}

	deadlineDate, err := time.Parse("2006-01-02 15:04:05", tb_Task.Tsk_deadline_date)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid input"})
		return
	}

	taskRepo := repositories.NewTaskRepository()
	task := models.Tb_Task{
		Tsk_name:          tb_Task.Tsk_name,
		Tsk_description:   tb_Task.Tsk_description,
		Tsk_creation_date: time.Now(),
		Tsk_update_date:   time.Now(),
		Tsk_deadline_date: deadlineDate,
		Tsk_color:         tb_Task.Tsk_color,
		Tskpr_id:          tb_Task.Tskpr_id,
		Tskst_id:          tb_Task.Tskst_id,
		Usr_id:            tb_Task.Usr_id,
	}

	var userId int

	userId, err = taskRepo.Insert(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to create task"})
		return
	}
	task.Tsk_id = userId

	c.JSON(http.StatusCreated, task)
}

// UpdateTask godoc
// @Summary Update an existing task
// @Description UpdateTask updates an existing task by ID
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param task_id path int true "Task ID"
// @Param task body models.Task_Request true "Task_Request"
// @Success 200 {object} models.Tb_Task
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /task/{task_id} [put]
func UpdateTask(c *gin.Context) {
	taskIDParam := c.Param("task_id")
	taskRepo := repositories.NewTaskRepository()

	taskID, err := strconv.Atoi(taskIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	var task_request models.Task_Request

	if err := c.ShouldBindJSON(&task_request); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid input"})
		return
	}

	deadlineDate, err := time.Parse("2006-01-02 15:04:05", task_request.Tsk_deadline_date)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid input"})
		return
	}

	task := models.Tb_Task{
		Tsk_id:            taskID,
		Tsk_name:          task_request.Tsk_name,
		Tsk_description:   task_request.Tsk_description,
		Tsk_creation_date: time.Now(),
		Tsk_update_date:   time.Now(),
		Tsk_deadline_date: deadlineDate,
		Tsk_color:         task_request.Tsk_color,
		Tskpr_id:          task_request.Tskpr_id,
		Tskst_id:          task_request.Tskst_id,
		Usr_id:            task_request.Usr_id,
	}

	if err := taskRepo.Update(&task); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// DeleteTask godoc
// @Summary Delete a task by ID
// @Description DeleteTask deletes a task by ID
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param task_id path int true "Task ID"
// @Success 204
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /task/{task_id} [delete]
func DeleteTask(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("task_id"))
	taskRepo := repositories.NewTaskRepository()
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	if err := taskRepo.Delete(taskID); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to delete task"})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetStatus godoc
// @Summary Get status by ID
// @Description GetStatus returns a status by ID
// @Tags status
// @Accept  json
// @Produce  json
// @Param tskst_id path int true "Status ID"
// @Success 200 {object} models.Tb_Task_status
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /status/{tskst_id} [get]
func GetStatus(c *gin.Context) {
	statusIDParam := c.Param("tskst_id")
	statusRepo := repositories.NewStatusRepository()

	statusID, err := strconv.Atoi(statusIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid status ID"})
		return
	}

	status, err := statusRepo.Select(statusID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, status)
}

// GetAllStatus godoc
// @Summary Get all status
// @Description GetAllStatus returns all status
// @Tags status
// @Accept  json
// @Produce  json
// @Success 200 {array} repositories.Task_status
// @Failure 500 {object} models.ErrorResponse
// @Router /status [get]
func GetAllStatus(c *gin.Context) {
	statusRepo := repositories.NewStatusRepository()
	allStatus, err := statusRepo.Select()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, allStatus)
}

// GetPriority godoc
// @Summary Get priority by ID
// @Description GetPriority returns a priority by ID
// @Tags priorities
// @Accept  json
// @Produce  json
// @Param tskpr_id path int true "Priority ID"
// @Success 200 {object} models.Tb_Task_priority
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /priority/{tskpr_id} [get]
func GetPriority(c *gin.Context) {
	priorityIDParam := c.Param("tskpr_id")
	priorityRepo := repositories.NewPriorityRepository()

	priorityID, err := strconv.Atoi(priorityIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid priority ID"})
		return
	}

	priority, err := priorityRepo.Select(priorityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, priority)
}

// GetPriorities godoc
// @Summary Get all priorities
// @Description GetPriorities returns all priorities
// @Tags priorities
// @Accept  json
// @Produce  json
// @Success 200 {array} repositories.Task_priority
// @Failure 500 {object} models.ErrorResponse
// @Router /priorities [get]
func GetPriorities(c *gin.Context) {
	priorityRepo := repositories.NewPriorityRepository()
	priorities, err := priorityRepo.Select()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, priorities)
}

// GetCategories godoc
// @Summary Get all categories
// @Description Get all categoreies from the database
// @Tags categories
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Tb_Category
// @Failure 500 {object} models.ErrorResponse
// @Router /categories [get]
func GetCategorys(c *gin.Context) {
	categorysRepo := repositories.NewCategoryRepository()
	categorys, err := categorysRepo.Select()

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, categorys)
}

// GetCategory godoc
// @Summary      Get category by ID
// @Description  Returns the details of a specific category by its ID
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        cat_id path int true "Category ID"
// @Success      200 {object} models.Tb_Category
// @Failure      400 {object} models.ErrorResponse
// @Failure      404 {object} models.ErrorResponse
// @Failure      500 {object} models.ErrorResponse
// @Router       /category/{cat_id} [get]
func GetCategory(c *gin.Context) {
	categoryIDParam := c.Param("cat_id")
	categoryRepo := repositories.NewCategoryRepository()

	categoryID, err := strconv.Atoi(categoryIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid category ID"})
		return
	}

	category, err := categoryRepo.Select(categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, category)
}

// CreateCategory godoc
// @Summary Create a new category
// @Description CreateCategory creates a new category in the system
// @Tags categories
// @Accept  json
// @Produce  json
// @Param category body models.Tb_Category true "Category data"
// @Success 201 {object} models.Tb_Category
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /category [post]
func CreateCategory(c *gin.Context) {
	var tb_Category models.Tb_Category
	if err := c.ShouldBindJSON(&tb_Category); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid input"})
		return
	}

	categoryRepo := repositories.NewCategoryRepository()
	category := models.Tb_Category{
		Cat_name: tb_Category.Cat_name,
		Usr_id:   tb_Category.Usr_id,
	}

	if err := categoryRepo.Insert(&category); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to create category"})
		return
	}
	c.JSON(http.StatusCreated, category)
}

// UpdateCategory godoc
// @Summary Update a category by ID
// @Description UpdateCategory updates a category by ID
// @Tags categories
// @Accept  json
// @Produce  json
// @Param cat_id path int true "Category ID"
// @Param category body models.Tb_Category true "Category data"
// @Success 200 {object} models.Tb_Category
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /category/{cat_id} [put]
func UpdateCategory(c *gin.Context) {
	categoryIDParam := c.Param("cat_id")
	categoryRepo := repositories.NewCategoryRepository()

	categoryID, err := strconv.Atoi(categoryIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid category ID"})
		return
	}

	var categoryDTO models.Category_DTO
	if err := c.ShouldBindJSON(&categoryDTO); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "invalid input"})
		return
	}

	category := models.Tb_Category{
		Cat_id:   categoryID,
		Cat_name: categoryDTO.Cat_name,
		Usr_id:   categoryDTO.Usr_id,
	}

	err = categoryRepo.Update(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	categoryResponse := models.Category_Response{
		Cat_id:   category.Cat_id,
		Cat_name: category.Cat_name,
		Usr_id:   category.Usr_id,
	}

	c.JSON(http.StatusOK, categoryResponse)
}

// DeleteCategory godoc
// @Summary Delete a category by ID
// @Description DeleteCategory deletes a category by its ID from the database
// @Tags categories
// @Accept  json
// @Produce  json
// @Param cat_id path int true "Category ID"
// @Success 200 {object} models.Tb_Category
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /category/{cat_id} [delete]
func DeleteCategory(c *gin.Context) {
	categoryIDParam := c.Param("cat_id")
	categoryRepo := repositories.NewCategoryRepository()

	categoryID, err := strconv.Atoi(categoryIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid category ID"})
		return
	}

	err = categoryRepo.Delete(categoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetTaskCategories godoc
// @Summary Get all task categories
// @Description Get all task categories from the database
// @Tags TaskCategory
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Tb_Task_category
// @Failure 500 {object} models.ErrorResponse
// @Router /taskcategories [get]
func GetTaskCategories(c *gin.Context) {
	taskCategorysRepo := repositories.NewTaskCategoryRepository()
	taskCategories, err := taskCategorysRepo.Select([]int{}, []int{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, taskCategories)
}

// GetTaskCategory godoc
// @Summary Get task category by task ID
// @Description Get a task category based on the task ID
// @Tags TaskCategory
// @Accept  json
// @Produce  json
// @Param tsk_id query int true "Task ID"
// @Success 200 {array} models.Tb_Task_category
// @Failure 400 {object} models.ErrorResponse "Invalid task ID"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /taskcategory/{tsk_id} [get]
func GetTaskCategory(c *gin.Context) {
	taskCategorysRepo := repositories.NewTaskCategoryRepository()

	tsk_id := c.Param("tsk_id")

	tskID, err := strconv.Atoi(tsk_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	taskCategories, err := taskCategorysRepo.Select([]int{tskID}, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	if len(taskCategories) == 0 {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Task Category not found"})
		return
	}

	c.JSON(http.StatusOK, taskCategories)
}

// CreateTaskCategory godoc
// @Summary Create a new task category
// @Description Create a new task category in the database
// @Tags TaskCategory
// @Accept  json
// @Produce  json
// @Param category body models.Tb_Task_category true "Task Category data"
// @Success 201 {object} models.Tb_Task_category
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /taskcategories [post]
func CreateTaskCategory(c *gin.Context) {
	taskCategoryRepo := repositories.NewTaskCategoryRepository()

	var newTaskCategory models.Tb_Task_category

	if err := c.ShouldBindJSON(&newTaskCategory); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid input data"})
		return
	}

	insertedID, err := taskCategoryRepo.Insert(&newTaskCategory)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to create task category"})
		return
	}

	newTaskCategory.Tsk_id = insertedID

	c.JSON(http.StatusCreated, newTaskCategory)
}

// UpdateTaskCategory godoc
// @Summary Update task category by task ID
// @Description Update the category of a task based on the task ID
// @Tags TaskCategory
// @Accept  json
// @Produce  json
// @Param tsk_id path int true "Task ID"
// @Param taskCategory body models.Tb_Task_category true "Task Category Data"
// @Success 200 {object} models.Tb_Task_category
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 404 {object} models.ErrorResponse "Task Category not found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /taskcategory/{tsk_id} [put]
func UpdateTaskCategory(c *gin.Context) {
	taskCategorysRepo := repositories.NewTaskCategoryRepository()

	tsk_id := c.Param("tsk_id")

	tskID, err := strconv.Atoi(tsk_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	var taskCategory models.Tb_Task_category

	if err := c.ShouldBindJSON(&taskCategory); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid input"})
		return
	}

	taskCategory.Tsk_id = tskID

	updated, err := taskCategorysRepo.Update(&taskCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	if !updated {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Task Category not found"})
		return
	}

	c.JSON(http.StatusOK, taskCategory)
}

// DeleteTaskCategory godoc
// @Summary Delete a task category by tsk_id
// @Description Delete a specific task category using the tsk_id
// @Tags TaskCategory
// @Accept  json
// @Produce  json
// @Param tsk_id path int true "Task Category ID"
// @Success 204 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /taskcategory/{tsk_id} [delete]
func DeleteTaskCategory(c *gin.Context) {
	taskCategoryIDParam := c.Param("tsk_id")
	taskCategoryRepo := repositories.NewTaskCategoryRepository()

	taskCategoryID, err := strconv.Atoi(taskCategoryIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid category ID"})
		return
	}

	err = taskCategoryRepo.Delete(taskCategoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	c.Status(http.StatusNoContent)
}
