package database

import (
    "context"
    "log"

    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

func Connect() {
    var err error
    databaseUrl := "postgres://username:password@localhost:5432/mydatabase"

    dbPool, err = pgxpool.New(context.Background(), databaseUrl)
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
