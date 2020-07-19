package config

import (
    "fmt"
    "os"
    "strconv"

    "github.com/joho/godotenv"
)

type DatabaseConfig struct {
    Host, User, Password, Dbname string
    Port int
}

var DbCfg DatabaseConfig

func init() {
    if err := godotenv.Load(); err != nil {
        fmt.Println("No .env file found")
        panic(err)
    }

    DbCfg.Host, _ = os.LookupEnv("HOST")
    portStr, _ := os.LookupEnv("PORT")
    DbCfg.Port, _ = strconv.Atoi(portStr)
    DbCfg.User, _ = os.LookupEnv("USER")
    DbCfg.Password, _ = os.LookupEnv("PASSWORD")
    DbCfg.Dbname, _ = os.LookupEnv("DBNAME")
}

func (D *DatabaseConfig) GetDbConfig() (string, int, string, string, string) {
    return DbCfg.Host, DbCfg.Port, DbCfg.User, DbCfg.Password, DbCfg.Dbname
}
