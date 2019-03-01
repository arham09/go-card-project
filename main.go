package main

import "fmt"

func main() {
	cards := newDeck()

	handCards, remainingCards := deal(cards, 5)

	// handCards.print()
	fmt.Println("My Hand")
	fmt.Println(handCards.toString())
	// remainingCards.print()
	fmt.Println("Dealer Hand")
	fmt.Println(remainingCards.toString())
}
