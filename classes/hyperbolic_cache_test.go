package classes

import (
	"fmt"
	"testing"
)

// TestHyperbolicCache tests the HyperbolicCache functionality.
func TestHyperbolicCache(t *testing.T) {
	cache := NewHyperbolicCache(3)

	// Test initial state
	result := cache.Get(1)
	expected := -1
	if result != expected {
		t.Errorf("Get(1): got %d, want %d", result, expected)
	}

	// Test Put and Get
	cache.Put(1, 1)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(1))
	fmt.Println(cache.cache[1])
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(4, 4)

	result = cache.Get(1)
	expected = 1
	if result != expected {
		t.Errorf("Get(1): got %d, want %d", result, expected)
	}

	result = cache.Get(2)
	expected = -1
	if result != expected {
		t.Errorf("Get(2): got %d, want %d", result, expected)
	}

	// After Put(3), item with key 2 should be evicted
	result = cache.Get(1)
	expected = 1
	if result != expected {
		t.Errorf("Get(2) after Put(3): got %d, want %d", result, expected)
	}

	// After Put(4), item with key 1 should be evicted
	result = cache.Get(1)
	expected = 1
	if result != expected {
		t.Errorf("Get(1) after Put(4): got %d, want %d", result, expected)
	}

	// Check other values
	result = cache.Get(3)
	expected = 3
	if result != expected {
		t.Errorf("Get(3): got %d, want %d", result, expected)
	}

	result = cache.Get(4)
	expected = 4
	if result != expected {
		t.Errorf("Get(4): got %d, want %d", result, expected)
	}
}
