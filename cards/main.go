package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// What was written before:
// var card string
// card = "Ace of Spades"
// var card string = "Ace of spades"
// card := "Ace of Spades"
// card = "Five of Diamonds"
// card := newCard()
// fmt.Println(card)
// fmt.Println(cards)
// for i, card := range cards {
// 	fmt.Println(i, card)
// }
// hand, remainingDeck := deal(cards, 5)

// 	hand.print()
// 	remainingDeck.print()

// cards.saveToFile("my_cards")
// cards := newDeckFromFile("my_cards")
