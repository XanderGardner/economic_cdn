package cache

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// HyperbolicCache represents a simple Hyperbolic cache.
type HyperbolicCache struct {
	capacity int
	used     int
	cache    map[string]*CacheItem
	stats    Stats
}

// CacheItem represents an item in the Hyperbolic cache.
type CacheItem struct {
	key             string
	value           []byte
	frequency, time int64
}

// NewHyperbolicCache creates a new Hyperbolic cache with the given capacity.
func NewHyperbolicCache(capacity int) *HyperbolicCache {
	return &HyperbolicCache{
		capacity: capacity,
		used:     0,
		cache:    make(map[string]*CacheItem),
		stats: Stats{
			Hits:   0,
			Misses: 0,
		},
	}
}

func (hc *HyperbolicCache) MaxStorage() int {
	return hc.capacity
}
func (hc *HyperbolicCache) RemainingStorage() int {
	return hc.capacity - hc.used
}

// Get retrieves the value associated with the key in the Hyperbolic cache.
func (hc *HyperbolicCache) Get(key string) (value []byte, ok bool) {
	if item, ok := hc.cache[key]; ok {
		hc.stats.Hits += 1
		item.frequency++
		return item.value, true
	}
	hc.stats.Misses += 1
	return nil, false
}

// Remove
func (hc *HyperbolicCache) Remove(key string) (value []byte, ok bool) {
	if item, ok := hc.cache[key]; ok {
		hc.used -= len(key) + len(item.value)
		delete(hc.cache, key)
		return item.value, true
	} else {
		return nil, false
	}
}

// Put inserts a key-value pair into the Hyperbolic cache.
func (hc *HyperbolicCache) Set(key string, value []byte) bool {
	if len(key)+len(value) > hc.capacity {
		return false
	}

	if item, ok := hc.cache[key]; ok {
		// Update existing item.
		item.value = value
		item.frequency++
	} else {
		// Add new item.
		item := &CacheItem{key, value, 1, time.Now().UnixNano()}
		hc.cache[key] = item
		hc.used += len(key) + len(value)
		// If the cache is full, evict the item with the lowest score.
		for len(hc.cache) > hc.capacity {
			hc.evict()
		}
	}

	return true
}

func randomSample(m map[string]*CacheItem, sampleSize int) map[string]*CacheItem {
	rand.Seed(time.Now().UnixNano())

	// Convert map keys to a slice
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	// Shuffle the keys
	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})

	result := make(map[string]*CacheItem)
	for i := 0; i < sampleSize && i < len(keys); i++ {
		result[keys[i]] = m[keys[i]]
	}

	return result
}

// evict removes the item with the lowest score from the Hyperbolic cache.
func (hc *HyperbolicCache) evict() {
	minScore := math.MaxFloat64
	var minKey string
	sample := randomSample(hc.cache, min(len(hc.cache), 64))

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
	hc.used -= len(minKey) + len(hc.cache[minKey].value)
	delete(hc.cache, minKey)
}
func (hc *HyperbolicCache) Len() int {
	return len(hc.cache)
}
func (hc *HyperbolicCache) Stats() *Stats {
	return &hc.stats
}
