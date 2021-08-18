package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
)

func CreateBucket() {
	db, _ := bolt.Open("borp/dat.db", 0600, nil)
	defer db.Close()
	db.Update(func(t *bolt.Tx) error {
		t.CreateBucketIfNotExists([]byte(Today()))
		return nil
	})
}

func AddOccurances(emote []byte, num int) {
	db, _ := bolt.Open("borp/dat.db", 0600, nil)
	defer db.Close()
	db.Update(func(t *bolt.Tx) error {
		b := t.Bucket([]byte(Today()))
		currentCount := b.Get(emote)
		currentCountInt, _ := strconv.Atoi(string(currentCount))
		nextCount := currentCountInt + num
		b.Put(emote, []byte(strconv.Itoa(nextCount)))

		return nil
	})
}

func Today() string {
	return time.Now().Format("01-02-2006")
}

func ExportToday(w http.ResponseWriter, req *http.Request) {
	db, _ := bolt.Open("borp/dat.db", 0600, nil)
	defer db.Close()

	_ = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(Today()))
		b.ForEach(func(k, v []byte) error {
			fmt.Println(string(k), string(v))
			return nil
		})
		return nil
	})
}
