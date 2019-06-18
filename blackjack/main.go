package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

const objective = 42

type player struct {
	name  string
	cards []card
	score int
}

func newPlayer(name string, cards []card) player {
	return player{
		name:  name,
		cards: cards,
		score: 0,
	}
}

var players = map[string]player{}

type card struct {
	name   string
	values []int
}

func newCard(name string, values []int) card {
	return card{
		name:   name,
		values: values,
	}
}

var cards = map[string]card{
	"Ace":   newCard("Ace", []int{1, 11}),
	"Two":   newCard("Two", []int{2}),
	"Three": newCard("Three", []int{3}),
	"Four":  newCard("Four", []int{4}),
	"Five":  newCard("cinco", []int{5}),
	"Six":   newCard("Six", []int{6}),
	"Seven": newCard("Seven", []int{7}),
	"Eight": newCard("Eight", []int{8}),
	"Nine":  newCard("Nine", []int{9}),
	"Ten":   newCard("Ten", []int{10}),
	"Jack":  newCard("Jack", []int{10}),
	"Queen": newCard("Queen", []int{10}),
	"King":  newCard("King", []int{10}),
}

type BlackjackSim interface {
	Hit(player string, card string)
	Winner() string
}

type BlackjackImpl struct {
	situation map[string][]string
}

func New(initialSituation map[string][]string) BlackjackSim {
	return BlackjackImpl{
		situation: initialSituation,
	}
}

func (bj BlackjackImpl) Hit(player string, card string) {
	bj.situation[player] = append(bj.situation[player], cards[card].name)
}

func (bj BlackjackImpl) Winner() string {
	for player, playerCards := range bj.situation {
		p := newPlayer(player, []card{})
		for _, playerCard := range playerCards {
			p.cards = append(p.cards, cards[playerCard])
			p.score += cards[playerCard].values[0]
		}
		players[player] = p
	}

	var playerSlice []player
	for _, player := range players {
		playerSlice = append(playerSlice, player)
	}

	sort.Slice(playerSlice, func(i, j int) bool {
		return playerSlice[i].score > playerSlice[j].score
	})

	return playerSlice[0].name
}

func main() {
	initialSituation := make(map[string][]string, 2)
	initialSituation["Coupier"] = []string{"Ace"}
	initialSituation["Jugador Uno"] = []string{"Ace", "Eight"}

	blackjackSim := New(initialSituation)

	/* for _, player := range players {
		for {
			Print("Hit or panic? ")
			decision, _ := reader().ReadString('\n')
			switch decision {
			case "hit":
				Println(player)
			case "panic":
				break
			}
		}
	} */

	Println(blackjackSim)
}

func reader() *bufio.Reader {
	return bufio.NewReader(os.Stdin)
}
