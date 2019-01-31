package main

func main()  {
	// cards := newDeck()

	// cards.saveToFile("my_cards")
	results := newDeckFromFile("my_cards")
	results.print()
	// hand, remianingCards := deal(cards, 5)
	
	// hand.print()
	// remianingCards.print()

}

