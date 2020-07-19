package main

import (
    "records/models"
    "fmt"
)

func main() {
    fmt.Println("I'm alive")
    var rec models.Record
    rec.Id = 1
    rec.Author = "Satanic Warmaster"
    rec.Title = "Opferblut"
    rec.Format = "LP"
    rec.Genre = "Black Metal"
    rec.Category = "Vinyl"
    rec.Publisher = "No Colours Records"
    rec.Year = 2003

    fmt.Println(rec.ToString())
}
