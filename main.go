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
)

var hub = newHub()

// var addr = flag.String("addr", "localhost:8081", "http service address")

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Print("env file not found, make sure this is on prod")
	}

	client := twitch.NewAnonymousClient()

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		count := strings.Count(strings.ToLower(message.Message), "borpa")
		if count != 0 {
			fmt.Printf("found %d borpa \n", count)
			hub.broadcast <- []byte(strconv.Itoa(count))
		}
	})

	print("Wtf")
	client.Join(os.Getenv("CHAN"))
	go func() {
		err := client.Connect()
		if err != nil {
			panic(err)
		}
	}()

	go hub.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		serveWs(hub, w, r)
	})

	err := http.ListenAndServe("localhost:8081", nil)
	if err != nil {
		panic(err)
	}
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

// func serveWs(pool *Pool, w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("WebSocket Endpoint Hit")
// 	conn, err := Upgrade(w, r)
// 	if err != nil {
// 		fmt.Fprintf(w, "%+v\n", err)
// 	}

// 	client := &Client{
// 		Conn: conn,
// 		Pool: pool,
// 	}

// 	pool.Register <- client
// 	go client.Read()
// }

// func setupRoutes() {
// 	// pool := NewPool()
// 	// pool.Start()

// 	fmt.Print("enmd setuip[")
// 	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		serveWs(pool, w, r)
// 	})
// }
