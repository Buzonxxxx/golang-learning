package main

import "fmt"

func main() {
	//slice
	cards := []string{"Ace of Spades", newCard()}
	fmt.Println(cards)
	
	//array
	pets := [2]string{"dog", "cat"}
	fmt.Println(pets)

	//append
	cards = append(cards, "Six of Spades")
	fmt.Println(cards)

	// for loop
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}