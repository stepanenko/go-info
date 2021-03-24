package main

func main() {
	cards := newDeck()

	// cards.print()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
	
	fmt.Println([]byte("Hello")) // type conversion, will print [72 101 108 108 111] <- byte slice
	fmt.Println([]string(hand))  // [Ace of Spades Two of Spades Three of Spades Four of Spades Five of Spades]
}
