package repositories

import (
	"context"
	"fmt"
	database "golang_api/internal"
	"golang_api/pkg/models"
	"log"

	"github.com/jackc/pgx/v5"
)

// TaskRepository representa o repositório para operações com tarefas
type TaskRepository struct{}
type User struct{}
type Category struct{}
type Task_category struct{}
type Task_status struct{}
type Task_priority struct{}

// NewTaskRepository cria uma nova instância do repositório de tarefas
func NewTaskRepository() *TaskRepository {
	return &TaskRepository{}
}

func NewUserRepository() *User {
	return &User{}
}

func NewStatusRepository() *Task_status {
	return &Task_status{}
}

func NewPriorityRepository() *Task_priority {
	return &Task_priority{}
}

// Insert, Update, Delete de usuários
func (r *User) Insert(user *models.Tb_User) (int, error) {
	db := database.GetDB()

	query := `
		INSERT INTO tb_user (
			usr_name, usr_email, usr_password
		) VALUES ($1, $2, $3)
		RETURNING usr_id
	`
	var userID int
	err := db.QueryRow(context.Background(), query, user.Usr_name, user.Usr_email, user.Usr_password).Scan(&userID)
	if err != nil {
		log.Printf("Erro ao inserir usuário: %v", err)
		return 0, err
	}

	log.Printf("Usuário inserido com sucesso %d\n", userID)
	return userID, nil
}

func (r *User) Update(user *models.Tb_User) error {
	db := database.GetDB()

	query := `
		UPDATE tb_user SET
			usr_name = $1,
			usr_email = $2,
			usr_password = $3
		WHERE usr_id = $4
	`

	_, err := db.Query(context.Background(), query, user.Usr_name, user.Usr_email, user.Usr_password, user.Usr_id)
	if err != nil {
		log.Printf("Erro ao atualizar os dados do usuário: %v", err)
		return err
	}

	log.Printf("Dados do usuário atualizados com sucesso %d\n", user.Usr_id)
	return nil
}

func (r *User) Delete(userID int) error {
	db := database.GetDB()

	query := `
		DELETE FROM tb_user
		WHERE usr_id = $1
	`

	_, err := db.Exec(context.Background(), query, userID)
	if err != nil {
		log.Printf("Erro ao deletar o usuário %v", err)
		return err
	}

	log.Printf("Usuário deletado com sucesso %d\n", userID)
	return nil
}

func (r *User) Select(userId ...int) ([]*models.UserResponse, error) {
	db := database.GetDB()

	var query string
	var rows pgx.Rows
	var err error

	if len(userId) > 0 {
		query = `SELECT usr_id, usr_name, usr_email FROM tb_user WHERE usr_id = $1`
		rows, err = db.Query(context.Background(), query, userId[0])
	} else {
		query = `SELECT usr_id, usr_name, usr_email FROM tb_user`
		rows, err = db.Query(context.Background(), query)
	}

	if err != nil {
		log.Printf("Erro ao selecionar o(s) usuário(s) %v", err)
		return nil, err
	}
	defer rows.Close()

	var users []*models.UserResponse
	for rows.Next() {
		var user models.UserResponse
		err := rows.Scan(
			&user.Usr_id,
			&user.Usr_name,
			&user.Usr_email,
		)
		if err != nil {
			log.Printf("Erro ao escanear o usuário %v", err)
			return nil, err
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erro ao iterar sobre as linhas %v", err)
		return nil, err
	}

	log.Printf("Usuário(s) selecionado(s) com sucesso %v", users)
	return users, nil
}

func (r *User) SelectByEmail(email string) (*models.Tb_User, error) {
	db := database.GetDB()

	query := `SELECT usr_id, usr_name, usr_email, usr_password from tb_user WHERE usr_email = $1`

	row := db.QueryRow(context.Background(), query, email)

	var user models.Tb_User

	err := row.Scan(&user.Usr_id, &user.Usr_name, &user.Usr_email, &user.Usr_password)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		log.Printf("error finding user: %v", err)
		return nil, err
	}

	return &user, nil
}

// Insert, Update, Delete de tarefas
func (r *TaskRepository) Insert(task *models.Tb_Task) (int, error) {
	db := database.GetDB()

	query := `
        INSERT INTO tb_task (
            tsk_name, tsk_description, tsk_creation_date, tsk_update_date,
            tsk_deadline_date, tsk_color, tskpr_id, tskst_id, usr_id
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        RETURNING tsk_id
    `
	var tskId int
	err := db.QueryRow(context.Background(), query, task.Tsk_name, task.Tsk_description, task.Tsk_creation_date, task.Tsk_update_date, task.Tsk_deadline_date, task.Tsk_color, task.Tskpr_id, task.Tskst_id, task.Usr_id).Scan(&tskId)
	if err != nil {
		log.Printf("Erro ao inserir tarefa: %v", err)
		return 0, err
	}

	log.Printf("Tarefa inserida com sucesso com ID %d\n", task.Tsk_id)
	return tskId, nil
}

func (r *TaskRepository) Update(task *models.Tb_Task) error {
	db := database.GetDB()

	query := `
		UPDATE tb_task SET
			tsk_name = $1,
			tsk_description = $2,
			tsk_update_date = $3,
			tsk_deadline_date = $4,
			tsk_color = $5,
			tskpr_id = $6,
			tskst_id = $7,
			usr_id = $8
		WHERE tsk_id = $9
	`

	_, err := db.Query(context.Background(), query, task.Tsk_name, task.Tsk_description, task.Tsk_update_date, task.Tsk_deadline_date, task.Tsk_color, task.Tskpr_id, task.Tskst_id, task.Usr_id, task.Tsk_id)
	if err != nil {
		log.Printf("Erro ao atualizar os dados da tarefa: %v", err)
		return err
	}

	log.Printf("Dados da tarefa atualizados com sucesso %d\n", task.Tsk_id)
	return nil
}

func (r *TaskRepository) Delete(taskID int) error {
	db := database.GetDB()

	query := `
		DELETE FROM tb_task
		WHERE tsk_id =$1
	`

	_, err := db.Exec(context.Background(), query, taskID)
	if err != nil {
		log.Printf("Erro ao deletar a tarefa %v", err)
		return err
	}

	log.Printf("Tarefa deletada com sucesso %d\n", taskID)
	return nil
}

func (r *TaskRepository) Select(taskId ...int) ([]*models.Tb_Task, error) {
	db := database.GetDB()

	var query string
	var rows pgx.Rows
	var err error

	if len(taskId) > 0 {
		query = `SELECT tsk_id, tsk_name, tsk_description, tsk_creation_date, tsk_update_date, tsk_deadline_date, tsk_color, tskpr_id, tskst_id, usr_id FROM tb_task WHERE tsk_id = $1`
		rows, err = db.Query(context.Background(), query, taskId[0])
	} else {
		query = `SELECT tsk_id, tsk_name, tsk_description, tsk_creation_date, tsk_update_date, tsk_deadline_date, tsk_color, tskpr_id, tskst_id, usr_id FROM tb_task`
		rows, err = db.Query(context.Background(), query)
	}

	if err != nil {
		log.Printf("Erro ao selecionar a tarefa %v", err)
		return nil, err
	}

	defer rows.Close()

	var tasks []*models.Tb_Task
	for rows.Next() {
		var task models.Tb_Task
		err := rows.Scan(
			&task.Tsk_id,
			&task.Tsk_name,
			&task.Tsk_description,
			&task.Tsk_creation_date,
			&task.Tsk_update_date,
			&task.Tsk_deadline_date,
			&task.Tsk_color,
			&task.Tskpr_id,
			&task.Tskst_id,
			&task.Usr_id,
		)
		if err != nil {
			log.Printf("Erro ao selecionar a tarefa %v", err)
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erro ao iterar sobre linhas %v", err)
		return nil, err
	}

	log.Printf("Tarefa selecionada com sucesso %v", tasks)
	return tasks, nil
}

// Insert, Update, Delete de Categorias

func (r *Category) Insert(category *models.Tb_Category) error {
	db := database.GetDB()

	query := `
		INSERT INTO tb_category (
			cat_name, usr_id
		) VALUES ($1, $2)
		RETURNING cat_id
	`
	_, err := db.Exec(context.Background(), query, category.Cat_name, category.Usr_id)
	if err != nil {
		log.Printf("Erro ao inserir categoria: %v", err)
		return err
	}

	log.Printf("Categoria inserido com sucesso %d\n", category.Cat_id)
	return nil
}

func (r *Category) Update(category *models.Tb_Category) error {
	db := database.GetDB()

	query := `
		UPDATE tb_category SET
			cat_name = $1,
			usr_id = $2,
		WHERE cat_id = $3
	`

	_, err := db.Exec(context.Background(), query, category.Cat_name, category.Usr_id)
	if err != nil {
		log.Printf("Erro ao atualizar os dados da categoria: %v", err)
		return err
	}

	log.Printf("Dados da categoria atualizados com sucesso %d\n", category.Cat_id)
	return nil
}

func (r *Category) Delete(category *models.Tb_Category) error {
	db := database.GetDB()

	query := `
		DELETE FROM tb_category
		WHERE cat_id =$1
	`

	_, err := db.Exec(context.Background(), query, category.Cat_id)
	if err != nil {
		log.Printf("Erro ao deletar a categoria %v", err)
		return err
	}

	log.Printf("Categoria deletada com sucesso %d\n", category.Cat_id)
	return nil
}

func (r *Category) Select(Cat_id int) (*models.Tb_Category, error) {
	db := database.GetDB()

	query := `SELECT cat_id, cat_name, usr_id FROM tb_category WHERE cat_id = $1`

	var category models.Tb_Category

	err := db.QueryRow(context.Background(), query, Cat_id).Scan(
		&category.Cat_id,
		&category.Cat_name,
		&category.Usr_id,
	)
	if err != nil {
		log.Printf("Erro ao selecionar a categoria %v", err)
		return nil, err
	}

	log.Printf("Categoria selecionada com sucesso %v", category)
	return &category, nil
}

// Insert, Update, Delete de tarefas e categorias

func (r *Task_category) Insert(task_category *models.Tb_Task_category) error {
	db := database.GetDB()

	query := `
		INSERT INTO tb_task_category (
			tsk_id, cat_id
		) VALUES ($1, $2)
		RETURNING tsk_id, cat_id
	`
	_, err := db.Exec(context.Background(), query, task_category.Tsk_id, task_category.Cat_id)
	if err != nil {
		log.Printf("Erro ao inserir tarefas e categorias: %v", err)
		return err
	}

	log.Printf("Tarefas e categorias inserido com sucesso %d\n", task_category.Tsk_id, task_category.Cat_id)
	return nil
}

func (r *Task_category) Update(task_category *models.Tb_Task_category) error {
	db := database.GetDB()

	query := `
		UPDATE tb_task_category SET
			tsk_id = $1,
			cat_id = $2,
		WHERE tsk_id = $3 and cat_id = $4
	`

	_, err := db.Exec(context.Background(), query, task_category.Tsk_id, task_category.Cat_id)
	if err != nil {
		log.Printf("Erro ao atualizar os dados das tarefas e categorias: %v", err)
		return err
	}

	log.Printf("Dados das tarefas e categorias atualizados com sucesso %d\n", task_category.Tsk_id, task_category.Cat_id)
	return nil
}

func (r *Task_category) Delete(task_category *models.Tb_Task_category) error {
	db := database.GetDB()

	query := `
		DELETE FROM tb_task_category
		WHERE tsk_id = $1 and cat_id = $2
	`

	_, err := db.Exec(context.Background(), query, task_category.Tsk_id, task_category.Cat_id)
	if err != nil {
		log.Printf("Erro ao deletar as  tarefas e categorias %v", err)
		return err
	}

	log.Printf("Tarefas e categorias deletadas com sucesso %d\n", task_category.Tsk_id, task_category.Cat_id)
	return nil
}

func (r *Task_category) Select(Tsk_id int, Cat_id int) (*models.Tb_Task_category, error) {
	db := database.GetDB()

	query := `SELECT tsk_id, cat_id FROM tb_task_category WHERE tsk_id = $1 and cat_id = $2`

	var task_category models.Tb_Task_category

	err := db.QueryRow(context.Background(), query, Cat_id).Scan(
		&task_category.Tsk_id,
		&task_category.Cat_id,
	)
	if err != nil {
		log.Printf("Erro ao selecionar as tarefas e categorias %v", err)
		return nil, err
	}

	log.Printf("Tarefas e categorias selecionadas com sucesso %v", task_category)
	return &task_category, nil
}

// Select de status

func (r *Task_status) Select(Tskst_id ...int) ([]*models.Tb_Task_status, error) {
	db := database.GetDB()

	var query string
	var rows pgx.Rows
	var err error

	if len(Tskst_id) > 0 {
		query = `SELECT tskst_id, tskst_name FROM tb_task_status WHERE tskst_id = $1`
		rows, err = db.Query(context.Background(), query, Tskst_id[0])
	} else {
		query = `SELECT tskst_id, tskst_name FROM tb_task_status`
		rows, err = db.Query(context.Background(), query)
	}

	if err != nil {
		log.Printf("Erro ao selecionar status %v", err)
		return nil, err
	}

	defer rows.Close()

	var all_task_status []*models.Tb_Task_status

	for rows.Next() {
		var task_status models.Tb_Task_status
		err := rows.Scan(
			&task_status.Tskst_id,
			&task_status.Tskst_name,
		)
		if err != nil {
			log.Printf("Erro ao selecionar a status %v", err)
			return nil, err
		}
		all_task_status = append(all_task_status, &task_status)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erro ao iterar sobre linhas %v", err)
		return nil, err
	}

	log.Printf("Status selecionado com sucesso %v", all_task_status)
	return all_task_status, nil
}

// Select de prioridade

func (r *Task_priority) Select(Tskpr_id ...int) ([]*models.Tb_Task_priority, error) {
	db := database.GetDB()

	var query string
	var rows pgx.Rows
	var err error

	if len(Tskpr_id) > 0 {
		query = `SELECT tskpr_id, tskpr_name FROM tb_task_priority WHERE tskpr_id = $1`
		rows, err = db.Query(context.Background(), query, Tskpr_id[0])
	} else {
		query = `SELECT tskpr_id, tskpr_name FROM tb_task_priority`
		rows, err = db.Query(context.Background(), query)
	}

	if err != nil {
		log.Printf("Erro ao selecionar prioridade %v", err)
		return nil, err
	}

	defer rows.Close()

	var task_priorities []*models.Tb_Task_priority

	for rows.Next() {
		var task_priority models.Tb_Task_priority
		err := rows.Scan(
			&task_priority.Tskpr_id,
			&task_priority.Tskpr_name,
		)
		if err != nil {
			log.Printf("Erro ao selecionar a prioridade %v", err)
			return nil, err
		}
		task_priorities = append(task_priorities, &task_priority)
	}
	if err = rows.Err(); err != nil {
		log.Printf("Erro ao iterar sobre linhas %v", err)
		return nil, err
	}

	log.Printf("Prioridade selecionada com sucesso %v", task_priorities)
	return task_priorities, nil
}
