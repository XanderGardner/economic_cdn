package cache

import (
	"container/list"
	// "fmt"
)

type NodeValue struct {
	Key string
	Value []byte
}

type FIFO struct {
	capacity int
	used int
	myMap map[string]*list.Element
	dll *list.List
	stats Stats
	// whatever fields you want here
}

// NewFIFO returns a pointer to a new FIFO with a capacity to store limit bytes
func NewFifo(limit int) *FIFO {
	fifo := FIFO{
		capacity: limit,
		used: 0,
		myMap: make(map[string]*list.Element),
		dll: list.New(),
		stats: Stats{
			Hits: 0,
			Misses: 0,
		},
	}
	return &fifo
}

// MaxStorage returns the maximum number of bytes this FIFO can store
func (fifo *FIFO) MaxStorage() int {
	return fifo.capacity
}

// RemainingStorage returns the number of unused bytes available in this FIFO
func (fifo *FIFO) RemainingStorage() int {
	return fifo.capacity - fifo.used
}

// Get returns the value associated with the given key, if it exists.
// ok is true if a value was found and false otherwise.
func (fifo *FIFO) Get(key string) (value []byte, ok bool) {
	node, exists := fifo.myMap[key]
	if exists{
		fifo.stats.Hits += 1
		return node.Value.(NodeValue).Value, true 
	} else {
		fifo.stats.Misses += 1
		return nil, false
	}
}

// Remove removes and returns the value associated with the given key, if it exists.
// ok is true if a value was found and false otherwise
func (fifo *FIFO) Remove(key string) (value []byte, ok bool) {
	node, exists := fifo.myMap[key]

	if exists{
		value := node.Value.(NodeValue).Value
		delete(fifo.myMap, key)
		// remove the node too
		fifo.used -= len(key) + len(value)
		fifo.dll.Remove(node)
		return value, true
	} else {
		return nil, false
	}
}

// Set associates the given value with the given key, possibly evicting values
// to make room. Returns true if the binding was added successfully, else false.
func (fifo *FIFO) Set(key string, value []byte) bool {
	// key value too big for entire capacity
	if len(key) + len(value) > fifo.capacity{
		return false
	}

	_, exists := fifo.myMap[key]
	new_node_value := NodeValue{
		Key: key,
		Value: value,
	}
	
	if exists {

		// remove end of queue until either there is room or the current key is at the end of the queue
		removed := false
		for true {
			if removed {
				if len(value) + len(key) + fifo.used <= fifo.capacity {
					break
				} else {
					lastnode := fifo.dll.Back()
					fifo.Remove(lastnode.Value.(NodeValue).Key)
				}

			} else {
				if len(value) - len(fifo.myMap[key].Value.(NodeValue).Value) + fifo.used <= fifo.capacity {
					break
				} else {
					// remove from end
					// if we removed the current key, set removed to true
					lastnode := fifo.dll.Back()
					if lastnode.Value.(NodeValue).Key == key {
						removed = true
					}
					fifo.Remove(lastnode.Value.(NodeValue).Key)
				}
			}
		}

		if !removed {
			// replace if we didn't remove
			old_node := fifo.myMap[key]
			new_node := fifo.dll.InsertAfter(new_node_value, old_node)
			fifo.used += len(value)+len(key)
			fifo.Remove(key)
			fifo.myMap[key] = new_node
		} else {
			// add to back if we did remove
			new_node := fifo.dll.PushBack(new_node_value)
			fifo.used += len(value)+len(key)
			fifo.myMap[key] = new_node
		}
	} else {
		// make room
		for fifo.capacity - fifo.used < len(value)+len(key) {
			// remove last node
			lastnode := fifo.dll.Back()
			fifo.Remove(lastnode.Value.(NodeValue).Key)
		} 

		// add to beginning of queue
		new_node := fifo.dll.PushFront(new_node_value)
		fifo.used += len(value)+len(key)
		fifo.myMap[key] = new_node
	}
	return true
}

// Len returns the number of bindings in the FIFO.
func (fifo *FIFO) Len() int {
	return len(fifo.myMap)
}

// Stats returns statistics about how many search hits and misses have occurred.
func (fifo *FIFO) Stats() *Stats {
	return &fifo.stats
}
