// main.go
package main

import (
	"fmt"
	"os"
	"strings"
	"io/ioutil"
	"unicode"
	"github.com/xander/economic_cdn/caches"
)






//////// helper functions

// Given a string, returns the strings with only characters in the alphabete and spaces
func keepAlphabeticAndSpaces(input string) string {
	var result []rune

	for _, char := range input {
		// Keep alphabetic characters and spaces
		if unicode.IsLetter(char) || unicode.IsSpace(char) {
			result = append(result, char)
		}
	}

	return string(result)
}

// for given keys and cache, returns the percent of cache hits from feeding all keys in
func testCache(message_text string, currCache cache.Cache) float64 {
	// Iterate over the words and request from cache
	words := strings.Split(message_text, " ")

	hit_count := 0.0
	miss_count := 0.0
	
	for _, word := range words {
		
		_, ok := currCache.Get(word)

		if ok {
			// hit
			hit_count += 1.0
			

		} else {
			// miss
			miss_count += 1.0

			// add to cache
			currCache.Set(word, []byte("test"))

		}
	}
	return hit_count / (hit_count + miss_count)
}


//////// main function

func main() {

	TEST_CACHE_SIZE := 10000
	
	if len(os.Args) <= 1 {
		return
	} 

	filePath := os.Args[1]
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return
	}
	message_text := keepAlphabeticAndSpaces(string(content))
	

	// test lru
	result := testCache(message_text, cache.NewLru(TEST_CACHE_SIZE))
	fmt.Printf("LRU Hit Rate: %f\n", result)

	// test fifo
	result = testCache(message_text, cache.NewFifo(TEST_CACHE_SIZE))
	fmt.Printf("Fifo Hit Rate: %f\n", result)
	
	// test hyperbolic
	result = testCache(message_text, cache.NewHyperbolicCache(TEST_CACHE_SIZE))
	fmt.Printf("Hyperbolic Hit Rate: %f\n", result)
	

	
}
