package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Deck create a new type of which is a slice of strings
type Deck []string

func (D Deck) print() {
	for index, card := range D {
		fmt.Println(index, card)
	}
}

func newDeck() Deck {
	cards := Deck{}
	cardSuits := []string{"Diamond", "Hearts", "Club", "Spades"}
	cardValues := []string{"Ace", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, value := range cardValues {
		for _, suit := range cardSuits {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(D Deck, handSize int) (Deck, Deck) {
	//dealHand :=
	//remainingHand :=

	return D[:handSize], D[handSize:]
}

func (D Deck) toString() string {
	return strings.Join(D, ",")
}

func (D Deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(D.toString()), 0666)

}

func (D Deck) deleteFile(filename string) error {
	return os.Remove(filename)
}

func newDeckFromFile(filename string) Deck {
	byteSlice, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return strings.Split(string(byteSlice), ",")
}

func (D Deck) shuffleDeck() Deck {

	rand.Shuffle(len(D), func(i, j int) {
		D[i], D[j] = D[j], D[i]
	})
	return D
}
