package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	fileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")

	timeLimit := flag.Int("time", 30, "time to answer the quiz in seconds")

	flag.Parse()

	file, err := os.Open(*fileName)

	if err != nil {
		fmt.Printf("Failed to open the csv file: %s \n", *fileName)
		os.Exit(1)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Failed to read the csv file")
	}

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct_aswers := 0
problemloop:
	for _, row := range records {

		fmt.Printf("%s:", row[0])

		answerCh := make(chan string)
		go func() {
			var user_answer string
			fmt.Scanf("%s\n", &user_answer)
			answerCh <- user_answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-answerCh:
			if answer == strings.TrimSpace(row[1]) {
				correct_aswers++
				fmt.Println("Correct answer!")
			} else {
				fmt.Println("Wrong answer :(")
			}
		}

	}

	fmt.Printf("Your score is %d out of %d \n", correct_aswers, len(records))

}
