package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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

// var hub = newHub()

var emoteCache = NewEmoteCache()

var dbc = NewDatabaseController()

var foundEmoteCache = make(map[string]int)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("env file not found, make sure this is on prod")
	}
	router := mux.NewRouter()

	client := twitch.NewAnonymousClient()

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		words := strings.Fields(message.Message)

		for _, word := range words {
			// count := 0
			_, ok := emoteCache.cache[word]
			if ok {
				foundEmoteCache[word] += 1
				// hub.broadcast <- []byte(strconv.Itoa(count))
			}
		}

		for emote, count := range foundEmoteCache {
			if count > 0 {
				log.Printf("Adding %d to %s", count, emote)
				emoteCache.cache[emote] += count
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

	// go hub.run()

	router.HandleFunc("/today", HandleToday).Methods("GET")
	router.HandleFunc("/history", HandleHistory).Methods("GET")

	router.HandleFunc("/health", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
	})

	http.Handle("/", router)

	// handler := cors.New(cors.Options{
	// 	AllowedOrigins: []string{"*"},
	// 	AllowedMethods: []string{"GET", "POST", "PATCH"},
	// 	AllowedHeaders: []string{"a_custom_header", "content_type"},
	// }).Handler(http.DefaultServeMux)

	res, err := http.Get("https://api.betterttv.net/3/cached/users/twitch/121059319")
	if err != nil {
		log.Fatal("coulnt get bttv res")
	}
	defer res.Body.Close()

	var dat BTTVUserResponse

	if err := json.NewDecoder(res.Body).Decode(&dat); err != nil {
		log.Fatal("json sucks")
	}

	emoteCache.Load(dat)
	dbc.CreateTodaysTable()

	StartCron()
	print("alldone")
	err = http.ListenAndServe(":8082", router)

	if err != nil {
		log.Fatal(err)
	}
}
