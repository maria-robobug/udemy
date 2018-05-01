package main

func main() {
	cards := deck{"Ace of Spades", newCard()}
	cards = append(cards, "Six of Spades")

	cards.printDeck()
}

func newCard() string {
	return "Five of Diamonds"
}
