package main

func main() {

	//var card = "Ace of Spades" // String is automatically assigned
	//card := newCard() // String is automatically assigned (infer)
	//cards := []string{};
	//fmt.Println(card)
	//cards := deck{"A", newCard()} // Slice
	//cards = append(cards, "Six of Spades")
	/*for i, card := range cards {
		fmt.Println(i, card)
	}*/
	//cards.print() // Print
	//fmt.Println(cards)

	cards := newDeck() //

	cards.print() // Print

}

/*func newCard() string {
	return "Five of Diamonds"
}*/
