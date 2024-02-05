package main

import (
	csv "encoding/csv"
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
	csvReader := csv.NewReader(file)
	questions, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(questions)
}
