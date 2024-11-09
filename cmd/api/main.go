package main

import (
	_ "golang_api/docs"
	"golang_api/internal"
	// "golang_api/pkg/models"
	"golang_api/pkg/repositories"
	"log"
	// "golang_api/pkg/models"
	// "golang_api/pkg/repositories"
	// "log"
	// "time"
)

// @title To-Do List Golang API
// @version 1.0
// @description This is a sample server.
// @host localhost:8080
// @BasePath /
func main() {
	// Conecta ao banco de dados
	database.Connect()
	defer database.Close()

	// Cria uma nova tarefa com os campos atualizados
	// task := &models.Tb_Task{
	// 	Tsk_name:          "Nova Tarefa",
	// 	Tsk_description:   "Descrição da tarefa",
	// 	Tsk_creation_date: time.Now(),
	// 	Tsk_update_date:   time.Now(),
	// 	Tsk_deadline_date: time.Now().Add(24 * time.Hour),
	// 	Tsk_color:         "Azul",
	// 	Tskpr_id:          7,
	// 	Tskst_id:          7,
	// }

	userId := 2
	
	repo := repositories.NewUserRepository()
	user, err := repo.Select(userId)
	if err != nil {
		log.Fatal("Erro ao inserir usuário: ", err)
	}

	log.Printf("Usuário encontrado: %v\n", user)

	// log.Printf("Usuário inserido com sucesso %d\n", task.Usr_id)

	// Cria o repositório de tarefas e insere a nova tarefa
	// repo := repositories.NewTaskRepository()
	// if err := repo.Insert(task); err != nil {
	// 	log.Fatalf("Erro ao inserir tarefa: %v\n", err)
	// }

	// log.Println("Tarefa inserida com sucesso:", task.Tsk_id)
}