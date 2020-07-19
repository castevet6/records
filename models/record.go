package models

import (
    "fmt"
    s "strings"
)

type Record struct {
    Id          int     `json:"id"`
    Author      string  `json:"author"`
    Title       string  `json:"title"`
    Format      string  `json:"format"`
    Category    string  `json:"category"`
    Genre       string  `json:"genre"`
    Publisher   string  `json:"publisher"`
    Year        int     `json:"year"`
}

func (r *Record) ToString() string {
    return fmt.Sprintf("%s: %s %s (%s, %d)", s.ToUpper(r.Author), r.Title, r.Format, r.Publisher, r.Year)
}
