package cache

import (
	"container/list"
	// "fmt"
)

// An LRU is a fixed-size in-memory cache with least-recently-used eviction
type LRU struct {
	capacity int
	used int
	myMap map[string]*list.Element
	dll *list.List
	stats Stats
	// whatever fields you want here
}

// NewLRU returns a pointer to a new LRU with a capacity to store limit bytes
func NewLru(limit int) *LRU {
	lru := LRU{
		capacity: limit,
		used: 0,
		myMap: make(map[string]*list.Element),
		dll: list.New(),
		stats: Stats{
			Hits: 0,
			Misses: 0,
		},
	}
	return &lru
}

// MaxStorage returns the maximum number of bytes this LRU can store
func (lru *LRU) MaxStorage() int {
	return lru.capacity
}

// RemainingStorage returns the number of unused bytes available in this LRU
func (lru *LRU) RemainingStorage() int {
	return lru.capacity - lru.used
}
func (lru *LRU) MoveToFront(key string, value []byte) {
	new_node_value := NodeValue{
		Key: key,
		Value: value,
	}
	lru.Remove(key)
	new_node := lru.dll.PushFront(new_node_value)
	lru.myMap[key] = new_node
	lru.used += len(value) + len(key)
}

// Get returns the value associated with the given key, if it exists.
// This operation counts as a "use" for that key-value pair
// ok is true if a value was found and false otherwise.
func (lru *LRU) Get(key string) (value []byte, ok bool) {
	node, exists := lru.myMap[key]
	if exists{
		lru.stats.Hits += 1
		val := node.Value.(NodeValue).Value
		lru.MoveToFront(key, val)
		return val, true 
	} else {
		lru.stats.Misses += 1
		return nil, false
	}
}

// Remove removes and returns the value associated with the given key, if it exists.
// ok is true if a value was found and false otherwise
func (lru *LRU) Remove(key string) (value []byte, ok bool) {
	node, exists := lru.myMap[key]

	if exists{
		value := node.Value.(NodeValue).Value
		delete(lru.myMap, key)
		// remove the node too
		lru.used -= len(key) + len(value)
		lru.dll.Remove(node)
		return value, true
	} else {
		return nil, false
	}
}

// Set associates the given value with the given key, possibly evicting values
// to make room. Returns true if the binding was added successfully, else false.
func (lru *LRU) Set(key string, value []byte) bool {
	// key value too big for entire capacity
	if len(key) + len(value) > lru.capacity{
		return false
	}

	_, exists := lru.myMap[key]
	new_node_value := NodeValue{
		Key: key,
		Value: value,
	}
	
	if exists {

		// remove end of queue until either there is room or the current key is at the end of the queue
		removed := false
		for true {
			if removed {
				if len(value) + len(key) + lru.used <= lru.capacity {
					break
				} else {
					lastnode := lru.dll.Back()
					lru.Remove(lastnode.Value.(NodeValue).Key)
				}

			} else {
				if len(value) - len(lru.myMap[key].Value.(NodeValue).Value) + lru.used <= lru.capacity {
					break
				} else {
					// remove from end
					// if we removed the current key, set removed to true
					lastnode := lru.dll.Back()
					if lastnode.Value.(NodeValue).Key == key {
						removed = true
					}
					lru.Remove(lastnode.Value.(NodeValue).Key)
				}
			}
		}

		if !removed {
			// replace if we didn't remove
			lru.Remove(key)
			new_node := lru.dll.PushFront(new_node_value)
			lru.used += len(value)+len(key)
			lru.myMap[key] = new_node
		} else {
			// add to back if we did remove
			new_node := lru.dll.PushFront(new_node_value)
			lru.used += len(value)+len(key)
			lru.myMap[key] = new_node
		}
	} else {
		// make room
		for lru.capacity - lru.used < len(value)+len(key) {
			// remove last node
			lastnode := lru.dll.Back()
			lru.Remove(lastnode.Value.(NodeValue).Key)
		} 

		// add to beginning of queue
		new_node := lru.dll.PushFront(new_node_value)
		lru.used += len(value)+len(key)
		lru.myMap[key] = new_node
	}
	return true
}

// Len returns the number of bindings in the LRU.
func (lru *LRU) Len() int {
	return len(lru.myMap)
}

// Stats returns statistics about how many search hits and misses have occurred.
func (lru *LRU) Stats() *Stats {
	return &lru.stats
}
