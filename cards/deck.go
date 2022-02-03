package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Creating type deck
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	var cardValues []string = []string{"Ace", "One", "Two", "Three", "Four", "Five", "Six", "Seven",
		"Eight", "Nine", "Ten", "Joker", "King", "Queen"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}

func (d deck) say() {

	for i, card := range d {
		fmt.Println(i, card)
	}

}

func deal(d deck, size int) (deck, deck) { return d[:size], d[size:] }

func (d deck) toString() string { return strings.Join([]string(d), ", ") }

func (d deck) saveToFile(file string) error {

	err := ioutil.WriteFile(file, []byte(d.toString()), 0666)
	return err

}

func deckFromFile(filename string) deck {

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		print("Something went wrong:")
		fmt.Println(err)
		os.Exit(1)
	}
	stringDeck := string(bytes)
	arr := strings.Split(stringDeck, ",")
	return deck(arr)

}
