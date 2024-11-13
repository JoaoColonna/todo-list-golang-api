#!/bin/sh

# Gerar a documentação Swagger
swag init -g ./cmd/api/main.go

# Iniciar a aplicação
./main