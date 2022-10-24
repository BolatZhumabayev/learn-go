package main

import (
	"fmt"
)

func main() {
	// cards := newDeck()

	// fmt.Println(cards.toString())
	// cards.saveToFile("./tests/my_cards")
	cards := newDeckFromFile("./tests/my_cards")
	cards.shuffle()
	cards.print()

}

func printState() {
	fmt.Println("State")
}
