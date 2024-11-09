package repositories

import (
	"context"
	"log"
	"golang_api/internal"
	"golang_api/pkg/models"
)

// TaskRepository representa o repositório para operações com tarefas
type TaskRepository struct{}
type User struct{}

// NewTaskRepository cria uma nova instância do repositório de tarefas
func NewTaskRepository() *TaskRepository {
    return &TaskRepository{}
}

func NewUserRepository() *User {
		return &User{}
}
//Insert, Update, Delete de usuários
func (r *User) Insert(user *models.Tb_User) error {
	db := database.GetDB()

	query := `
		INSERT INTO tb_user (
			usr_name, usr_email, usr_password
		) VALUES ($1, $2, $3)
		RETURNING usr_id
	`
	_, err := db.Exec(context.Background(), query, user.Usr_name, user.Usr_email, user.Usr_password)
	if err != nil {
		log.Printf("Erro ao inserir usuário: %v", err)
		return err
	}

	log.Printf("Usuário inserido com sucesso %d\n", user.Usr_id)
	return nil
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

	_, err := db.Exec(context.Background(), query, user.Usr_name, user.Usr_email, user.Usr_password, user.Usr_id)
	if err != nil {
		log.Printf("Erro ao atualizar os dados do usuário: %v", err)
		return err
	}

	log.Printf("Dados do usuário atualizados com sucesso %d\n", user.Usr_id)
	return nil
}

func (r *User) Delete(user *models.Tb_User) error {
	db := database.GetDB()

	query := `
		DELETE FROM tb_user
		WHERE usr_id =$1
	`

	_, err := db.Exec(context.Background(), query, user.Usr_id)
	if err != nil {
		log.Printf("Erro ao deletar o usuário %v", err)
		return err
	}

	log.Printf("Usuário deletado com sucesso %d\n", user.Usr_id)
	return nil
}

func (r *User) Select(userId int) (*models.Tb_User, error) {
	db := database.GetDB()

	query := `SELECT usr_id, usr_name, usr_email, usr_password FROM tb_user WHERE usr_id = $1`

	var user models.Tb_User

	err := db.QueryRow(context.Background(), query, userId).Scan(
		&user.Usr_id,
		&user.Usr_name,
		&user.Usr_email,
		&user.Usr_password,
	)
	if err != nil {
		log.Printf("Erro ao selecionar o usuário %v", err)
		return nil, err
	}

	log.Printf("Usuário selecionado com sucesso %v", user)
	return &user, nil
}

// Insert insere uma nova tarefa no banco de dados usando um modelo Tb_Task
func (r *TaskRepository) Insert(task *models.Tb_Task) error {
    db := database.GetDB()
    
    query := `
        INSERT INTO tb_task (
            tsk_name, tsk_description, tsk_creation_date, tsk_update_date,
            tsk_deadline_date, tsk_color, tskpr_id, tskst_id
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING tsk_id
    `
    
    _, err := db.Exec(context.Background(), query, task.Tsk_name, task.Tsk_description, task.Tsk_creation_date, task.Tsk_update_date, task.Tsk_deadline_date, task.Tsk_color, task.Tskpr_id, task.Tskst_id)
    if err != nil {
        log.Printf("Erro ao inserir tarefa: %v", err)
        return err
    }

    log.Printf("Tarefa inserida com sucesso com ID %d\n", task.Tsk_id)
    return nil
}