package main

import (
	_ "golang_api/docs"
	database "golang_api/internal"
	"os"

	// "golang_api/pkg/models"
	"golang_api/pkg/config"
	"golang_api/pkg/routes"

	// "golang_api/pkg/repositories"
	"log"

	ginSwagger "github.com/swaggo/gin-swagger"
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
	cfg := config.LoadConfig()

	// Connect to the database
	database.Connect()
	defer database.Close()

	swaggerHost := os.Getenv("SWAGGER_HOST")
	if swaggerHost == "" {
		swaggerHost = "localhost:8080"
	}
	ginSwagger.URL("http://" + swaggerHost + "/swagger/doc.json")

	r := routes.SetupRouter()

	log.Println("Iniciando servidor na porta:", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
