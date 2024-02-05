package main

import (
	csv "encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var correctAnswers int

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
		// Will exit if cell lengths are not all the same
		log.Fatal(err)
	}

	totalQuestions := len(questions)

	// Iterate through questions
	for _, question := range questions {
		// Ask questions
		fmt.Println(question[0])
		var answer string
		fmt.Scanln(&answer)

		if answer == question[1] {
			fmt.Println("Correct!")
			correctAnswers++
		} else {
			fmt.Println("Whoops, that's not correct!")
		}
	}

	fmt.Printf("You got %d out of %d questions correct! Thanks for playing!\n", correctAnswers, totalQuestions)
}
