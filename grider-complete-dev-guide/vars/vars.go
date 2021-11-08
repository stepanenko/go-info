package main

import "fmt"

func main() {
	// var card string = "Ace of Spades" // valid but unnecessary
	// card := "Ace of Spades" // use := only once when initialise a new variable
	cards := []string{newCard(), "Queen of Diamond", newCard()}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Ace of Spades"
}
