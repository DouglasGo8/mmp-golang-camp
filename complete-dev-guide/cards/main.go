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

	//cards := newDeck() //

	//cards.print() // Print

	//hand, remain := deal(cards, 5)

	//hand.print()   // Print
	//remain.print() // Print

	//fmt.Println(cards.toString())

	//cards.saveToFile("my_cards")

	//cards := newDeckFromFile("my_cards")

	//cards.print() // Print

	cards := newDeck()

	cards.shuffle()
	cards.print() // Print
}
