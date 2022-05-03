package main

import "fmt"

func main() {
	// cards := newDeck()
	cards := newDeckFromFile("cards.txt")

	hand, remainingCards := deal(cards, 4)

	hand.print()
	remainingCards.print()
	fmt.Println(cards.toString())
	// cards.saveToFile("cards.txt")

}
