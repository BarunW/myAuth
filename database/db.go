package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	_ "github.com/lib/pq"
)

type DB struct{
   sqlDb *sql.DB 
}

const (
    dbHost = "localhost"
    dbPort = "5432"
    dbUser = "postgres"
    dbPassword= "mysecretpassword"
    dbName = "postgres"
    maxRetry int = 3
)


func InitDB() (*DB, error) {
    var (
        postgresDb *sql.DB
        err error
    )

    for retries := 0; retries < maxRetry; retries++{
        postgresDb, err = sql.Open("postgres", "")        
        if err != nil{
            fmt.Println("Error while connecting to database: Retrying", retries )
            <-time.After(10 * time.Second)
            continue
        }
        break
    }

    if err != nil{
        panic(err)
    }

    if err := postgresDb.Ping(); err != nil{
        slog.Error("Error while pinging to db", "Details", err.Error())
        return nil, err
    }

    return &DB{
        sqlDb : postgresDb,
    }, nil


}
