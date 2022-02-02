package main

import "fmt"

func main() {

	cards := newDeck()

	// cards.say()

	// deck1, deck2 := deal(cards, 5)

	// print("deck 1 is")
	// deck1.say()

	// print("deck 2 is")
	// deck2.say()

	// print(cards.toString())

	cards.saveToFile("deck1.txt")

}

func print(s string) {
	fmt.Println(s)
}
