package main

func main() {

	cards := newDeckFromFile("my_text")

	cards.shuffle()

	cards.print()
	// cards := newDeck()

	// cards.saveToFile("my_text")

	// handCards, remainingCards := deal(cards, 5)

	// // handCards.print()
	// fmt.Println("My Hand")
	// fmt.Println(handCards.toString())
	// // remainingCards.print()
	// fmt.Println("Dealer Hand")
	// fmt.Println(remainingCards.toString())
}
