package main

import (
	"fmt"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func (d deck) printDeck() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
