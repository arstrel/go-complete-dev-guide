package main

func main() {
	// cards := newDeck()
	cards := newDeckFromFile("cards.txt")

	cards.shuffle()
	cards.print()

	// cards.saveToFile("cards.txt")

}
