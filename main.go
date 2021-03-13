package main

import "fmt"

func main() {
	// var card string = "Ace of Spades" // valid but unnecessary
	// card := "Ace of Spades" // use := only once when initialise a new variable
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Ace of Spades"
}
