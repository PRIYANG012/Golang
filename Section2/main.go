package main

// import "fmt"

func main() {

	// var card string="Ace of Spades"  // One way of defining var
	// card := "Ace of Spades" // compiler will automatically recoganize the variable type
	// card := "Five of Diamonds" //Error
	// card = "Five of Diamonds"  // correct

	// card := newcard()
	// cards := deck{"Ace of Diamonds", newcard()}
	// cards = append(cards, "six of Spades")

	// cards := newDeck()  // get a new deck

	// fmt.Println(cards)

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }


	// hand,remainingCards:=deal(cards,5)  // divining the deck after dealing
	// hand.print()
	// remainingCards.print()

	// fmt.Println(cards.toString()) //coverting deck to single string



	// cards.saveToFile("My_cards_File")   //saving to file 

	cards:=newDeckFromFile("My_cards_File")
	cards.print()
	
}

func newcard() string {
	return "Five of Diamonds"
}
