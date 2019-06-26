package main

import (
	"fmt"
	"log"

	board "github.com/aleixcam/gopher/board-kata"
	parser "github.com/aleixcam/gopher/board-kata/parser"
)

func main() {
	msgs, err := board.ReadInput("data/input.csv")
	if err != nil {
		log.Fatal(err)
	}

	for i, msg := range msgs {
		msgs[i] = parser.Parse(msg)
	}

	fmt.Println(msgs)

	//TODO: Print output into an html file

	fmt.Println("done!")
}
