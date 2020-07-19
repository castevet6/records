package main

import (
    "fmt"
    "records/db"

)

func main() {
    fmt.Println("I'm alive")
    db.InitDB()
    fmt.Println(db.Db)
}
