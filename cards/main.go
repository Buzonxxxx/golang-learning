package main

import "fmt"

func main()  {
	cards := []string {newCard(), "Six of Spades"}
	cards = append(cards, "Ave of Diamonds")
	for i, card := range cards{
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}