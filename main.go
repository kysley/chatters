package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

var hub = newHub()

// keeps track of the -daily- usage
var borpas = make(map[string]int)

// var addr = flag.String("addr", "localhost:8081", "http service address")

func main() {
	borpas["borpa"] = 0
	borpas["borpaspin"] = 0
	borpas["kachorpa"] = 0
	borpas["moon2spin"] = 0
	borpas["cum"] = 0
	catalog := make([]string, 0, len(borpas))
	for k := range borpas {
		catalog = append(catalog, k)
	}

	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("env file not found, make sure this is on prod")
	}

	client := twitch.NewAnonymousClient()

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		words := strings.Fields(message.Message)

		for _, word := range words {
			count := 0
			lower := strings.ToLower(word)

			_, ok := borpas[lower]
			if ok {
				fmt.Printf("found %s \n", lower)
				count++
			}
			if count > 0 {
				borpas[lower] += count
				hub.broadcast <- []byte(strconv.Itoa(count))
			}
		}
	})

	// client.Join(os.Getenv("CHAN"))
	client.Join("MOONMOON")

	go func() {
		err := client.Connect()
		if err != nil {
			panic(err)
		}
	}()

	go hub.run()

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	http.HandleFunc("/health", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/catalog", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte(strings.Join(catalog, ",")))
	})
	http.HandleFunc("/today", func(w http.ResponseWriter, r *http.Request) {
		database, _ := sql.Open("sqlite3", "data/borp/sql.db")
		defer database.Close()

		rows, _ := database.Query("SELECT name, count from totals")

		var name string
		var count int
		for rows.Next() {
			rows.Scan(&name, &count)
			w.Write([]byte(fmt.Sprintf("%s,%d", name, count)))
		}
	})

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH"},
		AllowedHeaders: []string{"a_custom_header", "content_type"},
	}).Handler(http.DefaultServeMux)

	ticker := time.NewTicker(60 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				// do stuff
				// WriteCache()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	path, e := os.Getwd()
	if e != nil {
		log.Println(e)
	}
	fmt.Println(path)

	PrepareDatabase()
	print("Wtf")
	err := http.ListenAndServe(":8081", handler)

	if err != nil {
		log.Fatal(err)
	}

	print("wtf2")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
