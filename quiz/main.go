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

func answer(csvReader *csv.Reader, totalChan chan bool, correctChan chan bool, timeoutChan chan bool) {
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			timeoutChan <- true
		} else if err != nil {
			panic(err)
		} else {
			fmt.Printf(questionTpl, line[0])

			var userAnswer int
			_, err := fmt.Scanf("%d", &userAnswer)
			if err != nil {
				panic(err)
			}

			answer, err := strconv.Atoi(line[1])
			if err != nil {
				panic(err)
			}

			totalChan <- true
			if userAnswer == answer {
				correctChan <- true
			}
		}
	}
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

	totalChan := make(chan bool)
	correctChan := make(chan bool)
	timeoutChan := make(chan bool)

	go answer(csvReader, totalChan, correctChan, timeoutChan)
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
