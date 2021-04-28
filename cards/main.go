package main

import "fmt"

func main() {
	// var card string = "Meu Card"
	card := "Meu Card"
	card = "Updating Card"
	card = test()

	fmt.Println(card)
}

func test() string {
	return "Test String"
}
