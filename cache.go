package main

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
