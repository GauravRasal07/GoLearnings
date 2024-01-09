package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

// Created new type "deck" which is a slice of string
type deck[] string

func newDeck() deck {
	cards := deck{}

	cardSuits := [] string {"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := [] string {"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, val := range cardValues {
			cards = append(cards, val + " of " + suit)
		}
	}

	return cards
}

// This is a connector function between deck and function printCards which behaves like a method in class
// (d deck): means every variable of type deck has now access to the function
//		d: the actual copy of deck will be available as d in the function refers to this/self
func (d deck) printCards() {
	for i, card:= range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) writeToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) (deck) {
	bs, error := ioutil.ReadFile(filename)

	if error != nil {
		fmt.Println("Error: ", error)
		os.Exit(1)
	}

	// stringData := 
	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffleCards() {
	lenOfDeck := len(d)
	for i := range d {
		newIndex := rand.Intn(lenOfDeck - 1)

		d[i], d[newIndex] = d[newIndex], d[i]
	}
}