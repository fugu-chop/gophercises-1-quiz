package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

type problem struct {
	q string
	a string
}

func main() {
	var correctAnswers int

	// Parse command line flags
	randomiseQuestionsPtr := flag.Bool("randomise", false, "whether the quiz questions are in random order")
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

	// Convert questions to struct as this reduces dependency on CSV format
	parsedQuestions := convertProblemFormat(questions)

	if err != nil {
		// Will exit if cell lengths are not all the same
		log.Fatal(err)
	}

	totalQuestions := len(parsedQuestions)

	// Print Instructions
	fmt.Printf("Welcome to the quiz! You have %d seconds to answer the entire quiz!", *timerLengthPtr)
	fmt.Scanln()

	// Start timer
	// Use pointer as the goroutine will otherwise capture the value
	// at time of passing the variable
	go func(correctAnswers *int) {
		time.Sleep(time.Duration(*timerLengthPtr) * time.Second)
		fmt.Println("Time's up!")
		fmt.Printf("You answered %d questions correctly out of %d total questions\n", *correctAnswers, totalQuestions)

		os.Exit(0)
	}(&correctAnswers)

	// Randomise depending on flag
	if *randomiseQuestionsPtr {
		parsedQuestions = randomiseQuestions(parsedQuestions)
	}

	// Iterate through questions
	for _, question := range parsedQuestions {
		// Ask questions
		fmt.Println(question.q)
		var answer string
		fmt.Scanln(&answer)

		// Assume that the answers are all lowercase
		if cleanAnswer(answer) == question.a {
			fmt.Println("Correct!")
			correctAnswers++
		} else {
			fmt.Println("Whoops, that's not correct!")
		}
	}

	fmt.Printf("You got %d out of %d questions correct! Thanks for playing!\n", correctAnswers, totalQuestions)
}

func cleanAnswer(answer string) string {
	lowercaseString := strings.ToLower(answer)
	return strings.TrimFunc(lowercaseString, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

func randomiseQuestions(questions []problem) []problem {
	var newQuestionsOrder = []problem{}
	questionOrder := rand.Perm(len(questions))
	for _, order := range questionOrder {
		newQuestionsOrder = append(newQuestionsOrder, questions[order])
	}

	return newQuestionsOrder
}

func convertProblemFormat(questions [][]string) []problem {
	parsedProblems := make([]problem, len(questions))

	for idx, question := range questions {
		parsedProblems[idx] = problem{
			q: question[0],
			a: question[1],
		}
	}

	return parsedProblems
}
