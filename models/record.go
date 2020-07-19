package models

import (
    "fmt"
    s "strings"
)

type Record struct {
    id          int     `json:"id"`
    author      string  `json:"author"`
    title       string  `json:"title"`
    format      string  `json:"format"`
    category    string  `json:"category"`
    genre       string  `json:"genre"`
    publisher   string  `json:"publisher"`
    year        int     `json:"year"`
}

func (r *Record) toString() string {
    return fmt.Sprintf("%s: %s %s (%s, %d)", s.ToUpper(r.author), r.title, r.format, r.publisher, r.year)
}
