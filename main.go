package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"unicode"
)

func main() {
	var correctAnswers int

	// Parse command line flags
	timerLengthPtr := flag.Int("timer", 30, "how long the quiz is available for")
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

	// Print Instructions
	fmt.Printf("Welcome to the quiz! You have %d seconds to answer the entire quiz!", *timerLengthPtr)
	fmt.Scanln()

	// Start timer
	go func(correctAnswers *int) {
		time.Sleep(time.Duration(*timerLengthPtr) * time.Second)
		fmt.Println("Time's up!")
		fmt.Printf("You answered %d questions correctly out of %d total questions\n", *correctAnswers, totalQuestions)

		os.Exit(1)
	}(&correctAnswers)

	// Iterate through questions
	for _, question := range questions {
		// Ask questions
		fmt.Println(question[0])
		var answer string
		fmt.Scanln(&answer)

		// Assume that the answers are all lowercase
		if cleanAnswer(answer) == question[1] {
			fmt.Println("Correct!")
			correctAnswers++
		} else {
			fmt.Println("Whoops, that's not correct!")
		}
	}

	fmt.Printf("You got %d out of %d questions correct! Thanks for playing!\n", &correctAnswers, totalQuestions)
}

func cleanAnswer(answer string) string {
	lowercaseString := strings.ToLower(answer)
	return strings.TrimFunc(lowercaseString, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}
