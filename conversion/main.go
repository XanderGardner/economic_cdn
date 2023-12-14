package conversion

import (
	"encoding/binary"
	"fmt"
)

func IntToBytes(n int) []byte {
	byteSlice := make([]byte, 8) // Assuming a 64-bit int
	binary.LittleEndian.PutUint64(byteSlice, uint64(n))
	return byteSlice
}

func BytesToInt(byteSlice []byte) int {
	return int(binary.LittleEndian.Uint64(byteSlice))
}


func main() {
	number := 60

	// Convert int to []byte
	byteRepresentation := IntToBytes(number)
	fmt.Printf("Integer: %d, Byte representation: %v\n", number, byteRepresentation)

	// Convert []byte back to int
	resultingInteger := BytesToInt(byteRepresentation)
	fmt.Printf("Byte representation: %v, Integer: %d\n", byteRepresentation, resultingInteger)
}