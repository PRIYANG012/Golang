package main



func main() {

	// var card string="Ace of Spades"  // One way of defining var
	// card := "Ace of Spades" // compiler will automatically recoganize the variable type
	// card := "Five of Diamonds" //Error
	// card = "Five of Diamonds"  // correct

	// card := newcard()
	// cards := deck{"Ace of Diamonds", newcard()}
	// cards = append(cards, "six of Spades")

	cards := newDeck()
	// fmt.Println(cards)

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	hand,remainingCards:=deal(cards,5)
	hand.print()
	remainingCards.print()
}

func newcard() string {
	return "Five of Diamonds"
}
