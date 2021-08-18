package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/boltdb/bolt"
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

var hub = newHub()

var borpas = make(map[string]int)

// var addr = flag.String("addr", "localhost:8081", "http service address")

func main() {
	borpas["borpa"] = 0
	borpas["borpaspin"] = 0
	borpas["kachorpa"] = 0
	borpas["moon2spin"] = 0
	borpas["cum"] = 0

	if err := godotenv.Load(".env"); err != nil {
		fmt.Print("env file not found, make sure this is on prod")
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
				borpas[word] += count
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
	http.HandleFunc("/today", func(w http.ResponseWriter, r *http.Request) {
		db, _ := bolt.Open("borp/dat.db", 0600, nil)
		defer db.Close()

		_ = db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(Today()))
			b.ForEach(func(k, v []byte) error {
				// fmt.Println(string(k), string(v))
				w.Write([]byte(string(k) + "," + string(v)))
				return nil
			})
			return nil
		})
	})

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH"},
		AllowedHeaders: []string{"a_custom_header", "content_type"},
	}).Handler(http.DefaultServeMux)

	ticker := time.NewTicker(600 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				// do stuff
				PurgeCache()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	// defer db.Close()
	borpas["borpa"] = 100
	PurgeCache()

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

func PurgeCache() {
	CreateBucket()
	for key, count := range borpas {
		if count > 0 {
			AddOccurances([]byte(key), count)
		}
	}
}
