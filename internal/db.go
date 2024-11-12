package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var dbPool *pgxpool.Pool

func Connect() {
	var err error

	// Carregar variáveis de ambiente do arquivo .env
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	// Parse a configuração do banco de dados
	config, err := pgxpool.ParseConfig(databaseUrl)

	if err != nil {
		log.Fatalf("Unable to parse database URL: %v\n", err)
	}

	// Desabilitar o cache de prepared statements
	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	// Cria o pool de conexões
	dbPool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	log.Println("Conexão com o banco de dados estabelecida!")
}

func GetDB() *pgxpool.Pool {
	return dbPool
}

func Close() {
	if dbPool != nil {
		dbPool.Close()
		log.Println("Conexão com o banco de dados fechada.")
	}
}
