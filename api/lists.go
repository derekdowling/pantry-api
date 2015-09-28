package api

import (
	"log"
	"net/http"
)

// List has a name and items
type List struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Items []Item `json:"items"`
}

// Item has a name and amount
type Item struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

// CreateList creates a new list
func CreateList(w http.ResponseWriter, r *http.Request) {
	log.Printf("r.body = %+v\n", r.Body)
}

func GetLists(w http.ResponseWriter, r *http.Request) {
	log.Printf("r.body = %+v\n", r.Body)
}
