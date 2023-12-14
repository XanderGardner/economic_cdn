package cache

import (
	//"fmt"
	"testing"
	//"time"
)

// TestHyperbolicCache tests the HyperbolicCache functionality.
func TestHyperbolicCache(t *testing.T) {
	// Set up the cache
	c := NewHyperbolicCache(100)

	key := "testKey"
	value := []byte("testValue")

	// Test Set and Get
	if !c.Set(key, value) {
		t.Error("Failed to set value in cache")
	}

	val, ok := c.Get(key)
	if !ok || string(val) != string(value) {
		t.Error("Failed to retrieve value from cache")
	}

	// Test Remove
	val, ok = c.Remove(key)
	if !ok || string(val) != string(value) {
		t.Error("Failed to remove value from cache")
	}

	// Ensure the key is no longer present in the cache
	if _, ok := c.Get(key); ok {
		t.Error("Key still exists in cache after removal")
	}

	// Test cache eviction
	c = NewHyperbolicCache(50)

	for i := 0; i < 10; i++ {
		k := string(rune(i + 65)) // 'A', 'B', 'C', ...
		v := make([]byte, 5)
		c.Set(k, v)
	}

	exceedKey := "exceedKey"
	exceedValue := make([]byte, 60)
	if c.Set(exceedKey, exceedValue) {
		t.Error("Item larger than cache capacity was added")
	}

	if c.Len() != 10 {
		t.Error("Cache did not reach maximum capacity")
	}

}
