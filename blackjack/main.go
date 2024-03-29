package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

const objective int = 42

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

func (p player) setCard(playerCard string) {
	if playerCard == "Ace" {
		p.cards = append(p.cards, cards[playerCard])
	} else {
		p.cards = append([]card{cards[playerCard]}, p.cards...)
	}

	players[p.name] = p
}

func (p player) calcScore() {
	score := 0
	for _, card := range p.cards {
		score = closestScore(score, card.values)
	}

	p.score = score
	players[p.name] = p
}

func closestScore(score int, values []int) int {
	posibilities := make([]int, len(values))
	for i, value := range values {
		posibilities[i] = objective - (score + value)
	}

	index := 0
	if len(values) > 1 {
		if posibilities[index] >= 0 {
			if posibilities[index+1] >= 0 {
				if posibilities[index] > posibilities[index+1] {
					index = index + 1
				}
			}
		} else {
			if posibilities[index+1] < 0 {
				if posibilities[index] < posibilities[index+1] {
					index = index + 1
				}
			} else {
				index = index + 1
			}
		}
	}

	return score + values[index]
}

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
	"Five":  newCard("Five", []int{5}),
	"Six":   newCard("Six", []int{6}),
	"Seven": newCard("Seven", []int{7}),
	"Eight": newCard("Eight", []int{8}),
	"Nine":  newCard("Nine", []int{9}),
	"Ten":   newCard("Ten", []int{10}),
	"Jack":  newCard("Jack", []int{10}),
	"Queen": newCard("Queen", []int{10}),
	"King":  newCard("King", []int{10}),
}

var deck = map[string]int{
	"Ace":   4,
	"Two":   4,
	"Three": 4,
	"Four":  4,
	"Five":  4,
	"Six":   4,
	"Seven": 4,
	"Eight": 4,
	"Nine":  4,
	"Ten":   4,
	"Jack":  4,
	"Queen": 4,
	"King":  4,
}

func getCardNames(cards []card) (names []string) {
	for _, card := range cards {
		names = append(names, card.name)
	}

	return
}

func getRandomCard() string {
	for card, left := range deck {
		if left > 0 {
			deck[card]--
			return card
		}
	}

	return ""
}

type BlackjackSim interface {
	Hit(player string, card string)
	Winner() string
}

type BlackjackImpl struct {
	situation map[string][]string
}

func New(initialSituation map[string][]string) BlackjackSim {
	for player := range initialSituation {
		players[player].calcScore()
	}

	return BlackjackImpl{
		situation: initialSituation,
	}
}

func (bj BlackjackImpl) Hit(player string, card string) {
	bj.situation[player] = append(bj.situation[player], cards[card].name)
	players[player].setCard(card)
	players[player].calcScore()
}

func (bj BlackjackImpl) Winner() string {
	for player := range bj.situation {
		Printf("%v scores %v points\n", player, players[player].score)
	}

	var playerSlice []player
	for _, player := range players {
		playerSlice = append(playerSlice, player)
	}

	sort.Slice(playerSlice, func(i, j int) bool {
		if playerSlice[i].score > objective {
			playerSlice[i].score = -1
		}

		if playerSlice[j].score > objective {
			playerSlice[j].score = -1
		}
		return playerSlice[i].score > playerSlice[j].score
	})

	if playerSlice[0].score == playerSlice[1].score {
		return ""
	}

	return playerSlice[0].name
}

func main() {
	players["Coupier"] = newPlayer("Coupier", []card{})
	players["Coupier"].setCard(getRandomCard())

	players["Jugador Uno"] = newPlayer("Jugador Uno", []card{})
	players["Jugador Uno"].setCard(getRandomCard())
	players["Jugador Uno"].setCard(getRandomCard())

	initialSituation := make(map[string][]string, 2)
	initialSituation[players["Coupier"].name] = getCardNames(players["Coupier"].cards)
	initialSituation[players["Jugador Uno"].name] = getCardNames(players["Jugador Uno"].cards)
	blackjackSim := New(initialSituation)

	Println("Initial situation:")
	for player, cards := range initialSituation {
		Printf("\t%v: %v\n", player, cards)
	}

	lost := false
	for player := range players {
		if player != "Coupier" {
			for {
				if players[player].score > objective {
					Println("Avobe limit")
					lost = true
					break
				}

				Printf("\n%v, you have %v points. Hit or panic? ", player, players[player].score)
				decision, _, _ := reader().ReadLine()
				if string(decision) == "hit" {
					card := getRandomCard()
					if card == "" {
						Println("No cards left in the deck")
						break
					}

					blackjackSim.Hit(player, card)
					Printf("%v hits a %v\n", player, card)
				} else if string(decision) == "panic" {
					Printf("%v backs off\n", player)
					break
				} else {
					Println("Wrong option")
				}
			}
		}
	}

	if !lost {
		Print("\n")
		for n := 0; n < 3; n++ {
			card := getRandomCard()
			if card == "" {
				Println("No cards left in the deck")
				break
			}

			blackjackSim.Hit("Coupier", card)
			Printf("%v hits a %v\n", "Coupier", card)
		}
		Printf("%v backs off\n", "Coupier")
	}

	Print("\n")
	winner := blackjackSim.Winner()
	if winner != "" {
		Printf("%v wins!\n", winner)
	} else {
		Println("Draw!")
	}
}

func reader() *bufio.Reader {
	return bufio.NewReader(os.Stdin)
}
