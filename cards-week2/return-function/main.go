package main

import "fmt"

func main() {
	card := "Ace of Spades"
	fmt.Println(card)
	
	fmt.Println(newCard())
}

func newCard() string {
	return "Five of Diamonds"
}