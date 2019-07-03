package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

const problemsFile = "problems.csv"
const questionTpl = "what %v, sir?"

var problems = make(map[int][]string)

func answer(totalChan chan bool, correctChan chan bool, timeoutChan chan bool) {
	for _, problem := range problems {
		fmt.Printf(questionTpl, problem[0])

		var userAnswer int
		_, err := fmt.Scanf("%d", &userAnswer)
		if err != nil {
			panic(err)
		}

		answer, err := strconv.Atoi(problem[1])
		if err != nil {
			panic(err)
		}

		totalChan <- true
		if userAnswer == answer {
			correctChan <- true
		}
	}

	timeoutChan <- true
}

func timeout(limit int, timeoutChan chan bool) {
	time.Sleep(time.Duration(limit) * time.Second)
	timeoutChan <- true
}

func main() {
	timeLimit := 10
	totalAnswers := 0
	correctAnswers := 0

	fileNamePtr := flag.String("file", problemsFile, "problems file name")
	flag.IntVar(&timeLimit, "tl", timeLimit, "time limit")

	flag.Parse()

	csvFile, err := os.Open(*fileNamePtr)
	if err != nil {
		log.Panicf("problems file not found: %v", *fileNamePtr)
	}

	csvReader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		} else {
			problems[len(problems)] = line
		}
	}

	totalChan := make(chan bool)
	correctChan := make(chan bool)
	timeoutChan := make(chan bool)

	go answer(totalChan, correctChan, timeoutChan)
	go timeout(timeLimit, timeoutChan)

	for {
		select {
		case <-totalChan:
			totalAnswers++
		case <-correctChan:
			correctAnswers++
		case <-timeoutChan:
			fmt.Printf("\n%v correct answers of %v\n", correctAnswers, totalAnswers)
			os.Exit(0)
		}
	}
}
