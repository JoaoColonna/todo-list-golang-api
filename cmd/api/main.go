package main

import (
    "log"
    "golang_api/pkg/config"
    "golang_api/pkg/routes"
)

func main() {
    cfg := config.LoadConfig()
    r := routes.SetupRouter()
    log.Println("Iniciando servidor na porta:", cfg.Port)
    if err := r.Run(":" + cfg.Port); err != nil {
        log.Fatal(err)
    }
}
