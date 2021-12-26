package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type DatabaseController struct {
	db *sql.DB
}

var database, _ = sql.Open("sqlite3", "chatters-data/sql.db")

func NewDatabaseController() *DatabaseController {
	return &DatabaseController{
		db: database,
	}
}

func Today() string {
	return time.Now().Format("01-02-2006")
}

func (c *DatabaseController) CreateTodaysTable() {
	_, e := c.db.Exec("CREATE TABLE IF NOT EXISTS '" + Today() + "' (" +
		"name TEXT," +
		"count	 INTEGER," +
		"PRIMARY KEY('name')" +
		")")

	if e != nil {
		log.Println("Error creating daily table" + e.Error())
	}
}

func (c *DatabaseController) PopulateRows(cache map[string]int) {
	for key := range cache {
		database.Exec(fmt.Sprintf(`INSERT OR IGNORE INTO '%s' (name, count) VALUES ($2, $3)`, Today()), key, 0)
	}
}

func (c *DatabaseController) AddEmoteOccurance(key string, count int) {
	log.Printf("Adding %d to %s", count, key)
	_, err := c.db.Exec(fmt.Sprintf(`UPDATE '%s' SET count = count + $2 WHERE name = $3`, Today()), count, key)

	if err != nil {
		log.Printf("ERROR Adding %d to %s", count, key)
		log.Print(err)
	}
}
