package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var database, _ = sql.Open("sqlite3", "borpa-data/sql.db")

func Today() string {
	return time.Now().Format("01-02-2006")
}

func WriteCache() {
	log.Print("Writing cache to database")
	_, e := database.Exec("CREATE TABLE IF NOT EXISTS '" + Today() + "' (" +
		"name	TEXT," +
		"count	INTEGER," +
		"PRIMARY KEY('name')" +
		")")

	statement, er := database.Prepare("INSERT into '" + Today() + "' (name, count) VALUES (?, ?)")
	for key := range borpas {
		log.Print(key)
		if er != nil {
			println(er.Error())
		}
		statement.Exec(key, 0)
	}
	statement.Close()

	if e != nil {
		log.Println("Error creating daily table" + e.Error())
	}

	for key, count := range borpas {
		if count > 0 {
			log.Printf("Writing emote: %s", key)
			query := fmt.Sprintf("UPDATE '%s' SET count = count + %d WHERE name = '%s'", Today(), count, key)
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

	_, e := database.Exec("CREATE TABLE IF NOT EXISTS totals ( " +
		"name	TEXT," +
		"count	INTEGER," +
		"PRIMARY KEY('name')" +
		")")

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
		statement.Exec(key, 0)
	}
	statement.Close()
}

func AddOccurance(key string, count int) {
	log.Printf("Adding %d to %s", count, key)
	stmt, _ := database.Prepare(`UPDATE totals SET count = count + ? WHERE name = ?`)

	_, err := stmt.Exec(count, key)

	if err != nil {
		log.Printf("ERROR Adding %d to %s", count, key)
		log.Print(err)
	}
}
