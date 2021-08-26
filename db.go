package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func Today() string {
	return time.Now().Format("01-02-2006")
}

func WriteCache() {
	log.Print("Writing cache to database")
	database, _ := sql.Open("sqlite3", "/data/sql.db")
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

func PrepareDatabase() {
	if _, err := os.Stat("borpa-data"); os.IsNotExist(err) {
		// path/to/whatever does not exist
		println("no data")
	}
	f, error := os.OpenFile("borpa-data/sql.db", os.O_CREATE, 0666)
	f.Close()
	if error != nil {
		log.Println("error creating db file", error)
	}
	database, error := sql.Open("sqlite3", "borpa-data/sql.db")
	if error != nil {
		log.Println(error)
	}
	_, e := database.Exec("CREATE TABLE IF NOT EXISTS totals ( " +
		"name	TEXT," +
		"count	INTEGER," +
		"PRIMARY KEY('name')" +
		")")
	// createStmt.Exec()
	// createStmt.Close()

	if e != nil {
		println(e)
	}
	log.Println("database opened!")
	log.Println(borpas)
	statement, er := database.Prepare("INSERT into totals (name, count) VALUES (?, ?)")
	for key := range borpas {
		log.Print(key)
		if er != nil {
			println(er.Error())
		}
		statement.Exec(key, 1)
	}
	statement.Close()
	database.Close()
}
