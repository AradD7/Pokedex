package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	memory 	map[string]cacheEntry
	mux 	*sync.RWMutex
}


type cacheEntry struct {
	createdAt 	time.Time
	val 		[]byte
}
