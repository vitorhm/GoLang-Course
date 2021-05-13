package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printBot(eb)
	printBot(sb)
}

func printBot(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hi, there"
}

func (sb spanishBot) getGreeting() string {
	return "Hola"
}
