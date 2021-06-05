package main

import "fmt"

func main() {

	cards := newDeck()

	hand, remainingCards := deal(cards, 5) // two return values, first assigned to hand variable and second assigned to remaingCards variable
	hand.print()
	remainingCards.print()

	// Join a slice of string
	fmt.Println(cards.tostring())

	//Save file to hard drive
	cards.saveToFile("my_cards")

	//Read file from hard drive
	//cards := newDeckFromFile("my_cards")
	//cards.print()

}
