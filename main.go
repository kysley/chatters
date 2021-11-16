package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

type BTTVUserResponse struct {
	ID            string        `json:"id"`
	Bots          []interface{} `json:"bots"`
	ChannelEmotes []struct {
		ID        string `json:"id"`
		Code      string `json:"code"`
		ImageType string `json:"imageType"`
		UserID    string `json:"userId"`
	} `json:"channelEmotes"`
	SharedEmotes []struct {
		ID        string `json:"id"`
		Code      string `json:"code"`
		ImageType string `json:"imageType"`
		User      struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			DisplayName string `json:"displayName"`
			ProviderID  string `json:"providerId"`
		} `json:"user"`
	} `json:"sharedEmotes"`
}

var hub = newHub()

var emoteCache = make(map[string]int)

var foundEmoteCache = make(map[string]int)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("env file not found, make sure this is on prod")
	}

	client := twitch.NewAnonymousClient()

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		words := strings.Fields(message.Message)

		for _, word := range words {
			// count := 0
			_, ok := emoteCache[word]
			if ok {
				foundEmoteCache[word] += 1
				// hub.broadcast <- []byte(strconv.Itoa(count))
			}
		}

		for emote, count := range foundEmoteCache {
			if count > 0 {
				// AddOccurance(emote, count)
				log.Printf("Adding %d to %s", count, emote)
				emoteCache[emote] += count
				foundEmoteCache[emote] = 0
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
	// http.HandleFunc("/catalog", func(rw http.ResponseWriter, r *http.Request) {
	// 	rw.Write([]byte(strings.Join(catalog, ",")))
	// })
	http.HandleFunc("/test", func(rw http.ResponseWriter, r *http.Request) {
		res, err := http.Get("https://api.betterttv.net/3/cached/users/twitch/121059319")
		if err != nil {
			log.Fatal("coulnt get bttv res")
		}
		defer res.Body.Close()

		var dat BTTVUserResponse

		if err := json.NewDecoder(res.Body).Decode(&dat); err != nil {
			log.Fatal("json sucks")
		}

		CacheLoad(dat)

		fmt.Print(dat.ChannelEmotes[0])
	})
	http.HandleFunc("/history", func(rw http.ResponseWriter, r *http.Request) {
		date := r.URL.Query().Get("day")

		var table string
		if date == "" {
			// table = "totals"
		} else {
			table = date
		}

		query := fmt.Sprintf(`SELECT name, count from '%s'`, table)
		rows, err := database.Query(query)

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("Don't have data for that date yo. Date format is ?day=mm-dd-yyyy"))
			return
		}

		var name string
		var count int
		for rows.Next() {
			rows.Scan(&name, &count)
			rw.Write([]byte(fmt.Sprintf("%s,%d \n", name, count)))
		}
	})

	http.HandleFunc("/today", func(rw http.ResponseWriter, r *http.Request) {
		for emote, count := range emoteCache {
			rw.Write([]byte(fmt.Sprintf("%s,%d \n", emote, count)))
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

	res, err := http.Get("https://api.betterttv.net/3/cached/users/twitch/121059319")
	if err != nil {
		log.Fatal("coulnt get bttv res")
	}
	defer res.Body.Close()

	var dat BTTVUserResponse

	if err := json.NewDecoder(res.Body).Decode(&dat); err != nil {
		log.Fatal("json sucks")
	}

	CacheLoad(dat)

	StartCron()
	print("alldone")
	err = http.ListenAndServe("localhost:85", handler)

	if err != nil {
		log.Fatal(err)
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
