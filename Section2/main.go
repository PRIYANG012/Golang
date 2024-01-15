package main

import "fmt"

func main() {

	// var card string="Ace of Spades"  // One way of defining var
	// card := "Ace of Spades" // compiler will automatically recoganize the variable type
	// card := "Five of Diamonds" //Error
	// card = "Five of Diamonds"  // correct

	// card := newcard()
	cards := []string{"Ace of Diamonds",newcard()}
	cards = append(cards,"six of Spades")

	// fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newcard() string{
	return "Five of Diamonds"
}