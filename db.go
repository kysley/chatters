package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func Today() string {
	return time.Now().Format("01-02-2006")
}

func WriteCache() {
	log.Print("Writing cache to database")
	database, _ := sql.Open("sqlite3", "borp/sql.db")
	defer database.Close()
	for key, count := range borpas {
		if count > 0 {
			log.Printf("Writing emote: %s", key)
			query := fmt.Sprintf("UPDATE totals SET count = count + %d WHERE name = '%s'", count, key)
			statement, error := database.Prepare(query)
			statement.Exec()

			if error != nil {
				log.Printf("Failed to update %s", key)
			}
		}
	}
	ResetCache()
}

func ResetCache() {
	log.Print("Clearing cache")

	for key := range borpas {
		borpas[key] = 0
	}
}

func test() {
	database, _ := sql.Open("sqlite3", "borp/sql.db")
	defer database.Close()
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS totals ( " +
		"name	TEXT," +
		"count	INTEGER," +
		"PRIMARY KEY('name')" +
		")")
	statement.Exec()

	for key := range borpas {
		statement, _ := database.Prepare("INSERT into totals (name, count) VALUES (?, ?)")
		statement.Exec(key, 0)
	}
}
