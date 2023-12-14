
package main

import (
	"encoding/binary"
	"fmt"
	"github.com/xander/economic_cdn/queue_caches"
)

func intToBytes(n int) []byte {
	byteSlice := make([]byte, 8) // Assuming a 64-bit int
	binary.LittleEndian.PutUint64(byteSlice, uint64(n))
	return byteSlice
}

func bytesToInt(byteSlice []byte) int {
	return int(binary.LittleEndian.Uint64(byteSlice))
}

func main() {
	number := 42

	// Convert int to []byte
	byteRepresentation := intToBytes(number)
	fmt.Printf("Integer: %d, Byte representation: %v\n", number, byteRepresentation)

	// Convert []byte back to int
	resultingInteger := bytesToInt(byteRepresentation)
	fmt.Printf("Byte representation: %v, Integer: %d\n", byteRepresentation, resultingInteger)

	c := cache.NewFifo(500)
	c.Set("hi", intToBytes(2))

	val, _ := c.Get("hi")
	fmt.Printf("Byte representation: %v, \n", val)
	

}
	