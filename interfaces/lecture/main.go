package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting_eb(eb)
	printGreeting_sb(sb)
}

func printGreeting_eb(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting_sb(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
