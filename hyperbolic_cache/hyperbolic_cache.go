package classes

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// HyperbolicCache represents a simple Hyperbolic cache.
type HyperbolicCache struct {
	capacity int
	cache    map[int]*CacheItem
}

// CacheItem represents an item in the Hyperbolic cache.
type CacheItem struct {
	key, value      int
	frequency, time int64
}

// NewHyperbolicCache creates a new Hyperbolic cache with the given capacity.
func NewHyperbolicCache(capacity int) *HyperbolicCache {
	return &HyperbolicCache{
		capacity: capacity,
		cache:    make(map[int]*CacheItem),
	}
}

// Get retrieves the value associated with the key in the Hyperbolic cache.
func (hc *HyperbolicCache) Get(key int) int {
	if item, ok := hc.cache[key]; ok {
		item.frequency++
		return item.value
	}
	return -1
}

// Put inserts a key-value pair into the Hyperbolic cache.
func (hc *HyperbolicCache) Put(key, value int) {
	if hc.capacity == 0 {
		return
	}

	if item, ok := hc.cache[key]; ok {
		// Update existing item.
		item.value = value
		item.frequency++
	} else {
		// Add new item.
		item := &CacheItem{key, value, 1, time.Now().UnixNano()}
		hc.cache[key] = item

		// If the cache is full, evict the item with the lowest score.
		if len(hc.cache) > hc.capacity {
			hc.evict()
		}
	}
}

func randomSample(m map[int]*CacheItem, sampleSize int) map[int]*CacheItem {
	rand.Seed(time.Now().UnixNano())

	// Convert map keys to a slice
	keys := make([]int, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	// Shuffle the keys
	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})

	result := make(map[int]*CacheItem)
	for i := 0; i < sampleSize && i < len(keys); i++ {
		result[keys[i]] = m[keys[i]]
	}

	return result
}

// evict removes the item with the lowest score from the Hyperbolic cache.
func (hc *HyperbolicCache) evict() {
	minScore := math.MaxFloat64
	var minKey int
	sample := randomSample(hc.cache, len(hc.cache))

	for key, item := range sample {
		fmt.Println(item.key)
		fmt.Println(time.Now().UnixNano() - item.time)
		score := float64(item.frequency) / float64(time.Now().UnixNano()-item.time)
		fmt.Println(score)
		if score < minScore {
			minScore = score
			minKey = key
		}
	}

	delete(hc.cache, minKey)
}
