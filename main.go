package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/cespare/xxhash/v2"
)

func changeEndianUint64(value uint64) uint64 {
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, value)
	return binary.BigEndian.Uint64(bs)
}

func main() {
	// Ensure a file path is provided as a command-line argument
	if len(os.Args) < 2 {
		log.Fatal("Please provide the file path as a command-line argument.")
	}

	// Get the file path from command-line argument
	filePath := os.Args[1]

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Create the xxHash64 hash instance
	hasher := xxhash.New()

	// Calculate the hash of the file
	if _, err := io.Copy(hasher, file); err != nil {
		log.Fatalf("Failed to calculate hash: %v", err)
	}

	// Get the resulting hash as a 64-bit unsigned integer
	hashResult := hasher.Sum64()

	// Print the hash result
	fmt.Printf("xxHash64 of %s: %016x\n", filePath, changeEndianUint64(hashResult))
}
