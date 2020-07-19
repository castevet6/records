package db

import (
    "database/sql"
    "fmt"

    "records/config"
    _ "github.com/lib/pq"
)

var Db *sql.DB

func InitDB() {
    var err error
    host, port, user, password, dbname := config.DbCfg.GetDbConfig()
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
    Db, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }

    err = Db.Ping()
    if err != nil {
        panic(err)
    }
    defer Db.Close()
}
