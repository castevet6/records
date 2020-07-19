package main

import (
    models "records/models"
    "fmt"
)

func main() {
    fmt.Println("I'm alive")
    r := models.Record{id:1, author:"Satanic Warmaster"}
    /* r.id = 1
    r.author = "Satanic Warmaster"
    r.title = "Opferblut"
    r.format = "LP"
    r.category = "Vinyl"
    r.genre = "Black Metal"
    r.publisher = "No Colours Records"
    r.year = 2003*/
}
