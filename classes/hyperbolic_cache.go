package main

import (
	"math"
)

// HyperbolicCache represents a simple Hyperbolic cache.
type HyperbolicCache struct {
	capacity int
	cache    map[int]*CacheItem
}

// CacheItem represents an item in the Hyperbolic cache.
type CacheItem struct {
	key, value, frequency int
	score                float64
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
		item.score = hyperbolicScore(item.frequency)
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
		item.score = hyperbolicScore(item.frequency)
	} else {
		// Add new item.
		item := &CacheItem{key, value, 1, hyperbolicScore(1)}
		hc.cache[key] = item

		// If the cache is full, evict the item with the lowest score.
		if len(hc.cache) > hc.capacity {
			hc.evict()
		}
	}
}

// evict removes the item with the lowest score from the Hyperbolic cache.
func (hc *HyperbolicCache) evict() {
	minScore := math.MaxFloat64
	var minKey int

	for key, item := range hc.cache {
		if item.score < minScore {
			minScore = item.score
			minKey = key
		}
	}

	delete(hc.cache, minKey)
}

// hyperbolicScore computes the hyperbolic score based on frequency.
func hyperbolicScore(frequency int) float64 {
	// Hyperbolic scoring function: score = 1 / (frequency + 1).
	return 1.0 / float64(frequency+1)
}
