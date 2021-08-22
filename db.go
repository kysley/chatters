package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func Today() string {
	return time.Now().Format("01-02-2006")
}

func WriteCache() {
	database, _ := sql.Open("sqlite3", "borp/sql.db")
	defer database.Close()
	for key, count := range borpas {
		if count > 0 {
			statement, _ := database.Prepare(fmt.Sprintf("UPDATE totals SET count = count + %d WHERE name = %s", count, key))
			statement.Exec()
		}
	}
}

func ResetCache() {
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
