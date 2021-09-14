package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

var hub = newHub()

// keeps track of the -daily- usage
var borpas = make(map[string]int)

var foundBorpas = make(map[string]int)

func main() {
	borpas["borpa"] = 0
	borpas["borpaspin"] = 0
	borpas["kachorpa"] = 0
	borpas["moon2spin"] = 0
	borpas["cum"] = 0
	borpas["corpa"] = 0
	borpas["borpau"] = 0
	borpas["cumdetected"] = 0
	borpas["batchest"] = 0
	borpas["batchesting"] = 0

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
				foundBorpas[lower] += 1
				hub.broadcast <- []byte(strconv.Itoa(count))
			}
		}

		for emote, count := range foundBorpas {
			if count > 0 {
				AddOccurance(emote, count)
				borpas[emote] += count
				foundBorpas[emote] = 0
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
	http.HandleFunc("/live", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	http.HandleFunc("/health", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/catalog", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte(strings.Join(catalog, ",")))
	})
	http.HandleFunc("/history", func(rw http.ResponseWriter, r *http.Request) {
		date := r.URL.Query().Get("day")

		var table string
		if date == "" {
			table = "totals"
		} else {
			table = date
		}

		query := fmt.Sprintf(`SELECT name, count from '%s'`, table)
		rows, err := database.Query(query)

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("Don't have data for that date yo. Date format is ?day=dd-mm-yyyy"))
			return
		}

		var name string
		var count int
		for rows.Next() {
			rows.Scan(&name, &count)
			rw.Write([]byte(fmt.Sprintf("%s,%d \n", name, count)))
		}
	})
	http.HandleFunc("/today", func(w http.ResponseWriter, r *http.Request) {
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

	path, e := os.Getwd()
	if e != nil {
		log.Println(e)
	}
	fmt.Println(path)

	PrepareDatabase()
	StartCron()
	print("Wtf")
	err := http.ListenAndServe(":80", handler)

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
