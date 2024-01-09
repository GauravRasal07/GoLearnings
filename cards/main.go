package main

import "fmt"

// var card string

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card = "King"

	// card := newCard()
	// fmt.Println(card)

	// printState()

	//Array: Fixed sized list
	//Slice: Variable sized list

	// cards := [] string {"Ace of spades", "Ace of diamonds", "Ace of Heart"}		//Declare a slice

	// cards = append(cards, "Seven of Spades")									//Appending to an existing slice

	// for i, card := range cards {												//Iterating through a slice
	// 	fmt.Println(i, card)
	// }

	// As we have declared deck as slice of string below initialization is also correct
	// cards := deck{"Ace of spades", "Ace of diamonds", "Ace of Heart"}

	cards := newDeck()

	// cards.printCards()

	// hand, remainingCards := deal(cards, 6)

	// fmt.Println("The cards in hand are: ")
	// hand.printCards()
	// fmt.Println("\nThe remaining cards are: ")
	// remainingCards.printCards()
	
	// fmt.Println(cards.toString())

	// cards.writeToFile("cards.txt")

	// newDeck := newDeckFromFile("cards.txt")

	cards.shuffleCards()

	fmt.Println("The content from file is: \n", cards)

	cards.printCards()
}