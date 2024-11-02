package main

import (
    "log"
    "golang_api/pkg/config"
    "golang_api/pkg/routes"
    _ "golang_api/docs"
)

// @title To-Do List Golang API
// @version 1.0
// @description This is a sample server.
// @host localhost:8080
// @BasePath /
func main() {
    cfg := config.LoadConfig()
    r := routes.SetupRouter()
    log.Println("Iniciando servidor na porta:", cfg.Port)
    if err := r.Run(":" + cfg.Port); err != nil {
        log.Fatal(err)
    }
}