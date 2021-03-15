package main

import "fmt"

// create a new type of 'deck' which is a slice of strings
// it will be kind of like a class in JS
type deck []string

// kind of a method of a deck 'class'
func (d deck) print() {
	// (d deck) - is a receiver of a function, 'd' here is like 'this' in JS
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five"}

	for _, s := range cardSuits {
		for _, n := range cardValues {
			cards = append(cards, n+" of "+s)

		}
	}

	return cards
}
