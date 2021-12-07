package main

import (
	"fmt"
	"log"
)

type EmoteCache struct {
	cache map[string]int
}

func NewEmoteCache() *EmoteCache {
	return &EmoteCache{
		cache: make(map[string]int),
	}
}

func (c *EmoteCache) Reset() {
	for key := range c.cache {
		c.cache[key] = 0
	}
}

func (c *EmoteCache) Load(dat BTTVUserResponse) {
	for _, val := range dat.ChannelEmotes {
		_, ok := emoteCache.cache[val.Code]

		if !ok {
			emoteCache.cache[val.Code] = 0
		}
	}

	for _, val := range dat.SharedEmotes {
		_, ok := emoteCache.cache[val.Code]

		if !ok {
			emoteCache.cache[val.Code] = 0
		}

	}
}

func (c *EmoteCache) Write() {
	log.Print("Writing cache to database")
	_, e := database.Exec("CREATE TABLE IF NOT EXISTS '" + Today() + "' (" +
		"name TEXT," +
		"count INTEGER," +
		"PRIMARY KEY('name')" +
		")")
	for key := range emoteCache.cache {
		_, er := database.Exec("INSERT into '"+Today()+"' (name, count) VALUES ($1, $2) returning count", key, 0)
		if er != nil {
			println(er.Error())
		}
	}

	if e != nil {
		log.Println("Error creating daily table" + e.Error())
	}

	for key, count := range emoteCache.cache {
		if count > 0 {
			log.Printf("Writing emote: %s", key)
			query := fmt.Sprintf("UPDATE '%s' SET count = count + %d WHERE name = '%s'", Today(), count, key)
			_, error := database.Exec(query)

			if error != nil {
				log.Printf("Failed to update %s", key)
			}
		}
	}
	c.Reset()
}
