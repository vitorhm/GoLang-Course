package main

func main() {
	// var card string = "Meu Card"
	// card := "Meu Card"
	// card = "Updating Card"
	// card = test()

	//cards := deck{"Teste", "Teste2"}

	cards := newDeck()

	hand, remaining := deal(cards, 5)

	hand.print()
	println("")
	remaining.print()
}
