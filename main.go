package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// Parse command line flags
	fileNamePtr := flag.String("filename", "./problems.csv", "file location")
	flag.Parse()

	// Open file
	file, err := os.Open(*fileNamePtr)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read data from file into memory
	buffer := make([]byte, 1024)
	rows, err := file.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", rows, buffer[:rows])
}
