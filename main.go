package main

import (
	"fmt"
	"log"
	"os"
)

// Need to read in a file (assume in the same director)
// Defaults to problems.csv, but need to accept other options via flag(?)
// Quizzes are short, so can read entire file into memory
func main() {
	// Open file
	file, err := os.Open("./problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read data from file into memory
	buffer := make([]byte, 2048)
	rows, err := file.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", rows, buffer[:rows])
}
