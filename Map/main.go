package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	// add
	colors["GG"] = "XDDD"

	// delete 
	delete(colors, "white")

	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

	// Declare 
	// type1
	// var colors map[string]string

	// type2
	// colors := make(map[string]string)

	// type3
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }