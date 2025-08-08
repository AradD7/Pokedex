package pokeapi

import (
	"net/http"
	"time"

	"github.com/AradD7/Pokedex/internal/pokecache"
)




type Clinet struct {
	memory 		pokecache.Cache
	httpClient 	http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Clinet {
	return Clinet{
		memory: 	pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}


