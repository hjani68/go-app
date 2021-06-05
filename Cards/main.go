package main

func main() {

	cards := newDeck()

	hand, remainingCards := deal(cards, 5) // two return values, first assigned to hand variable and second assigned to remaingCards variable
	hand.print()
	remainingCards.print()

}
