package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	client := twitch.NewAnonymousClient()

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Printf("found %d borpa \n", strings.Count(strings.ToLower(message.Message), "borpa"))
	})

	client.Join(os.Getenv("CHAN"))

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}
