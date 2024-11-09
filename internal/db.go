package database

import (
    "context"
    "log"
    "github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

func Connect() {
    var err error
    databaseUrl := "postgresql://postgres.qbyxzucfcjuvzyeahthc:todolist0211A!@aws-0-sa-east-1.pooler.supabase.com:6543/postgres"

    // Parse a configuração do banco de dados
    config, err := pgxpool.ParseConfig(databaseUrl)
    if err != nil {
        log.Fatalf("Unable to parse database URL: %v\n", err)
    }

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