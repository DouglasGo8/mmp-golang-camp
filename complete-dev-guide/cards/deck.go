package main

import "fmt"

// Create a new type of 'deck' which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardsOfSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardOfValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardsOfSuits {
		for _, value := range cardOfValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}

// Receiver function
func (d deck) print() { // deck works like this Java keyword
	for i, card := range d {
		fmt.Println(i, card)
	}
}
