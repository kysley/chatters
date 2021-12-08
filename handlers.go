package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Emote struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type TodayResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func HandleToday(rw http.ResponseWriter, r *http.Request) {
	rows, err := database.Query(fmt.Sprintf("SELECT * FROM '%s'", Today()))

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var emotes []Emote

	var name string
	var count int
	for rows.Next() {
		rows.Scan(&name, &count)
		emotes = append(emotes, Emote{name, count})
	}
	json.NewEncoder(rw).Encode(emotes)
	rw.Header().Add("Content-Type", "application/json")
}

func HandleHistory(rw http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")

	table := date

	query := fmt.Sprintf(`SELECT name, count from '%s'`, table)
	rows, err := database.Query(query)

	var emotes []Emote

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Don't have data for that date yo. Date format is ?date=mm-dd-yyyy"))
		return
	}
	defer rows.Close()

	var name string
	var count int
	for rows.Next() {
		rows.Scan(&name, &count)
		emotes = append(emotes, Emote{name, count})
	}

	json.NewEncoder(rw).Encode(emotes)
	rw.Header().Add("Content-Type", "application/json")
}
