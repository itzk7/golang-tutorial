package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

// Receiver functionsf
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() deck {
	cardSuits := []string{"Diamond", "Spade", "Clover", "Heart"}
	cardValues := []string{"1", "2", "3"}
	cards := deck{}
	for _, suit := range cardSuits {
		for _, card := range cardValues {
			cards = append(cards, card+" of "+suit)
		}
	}
	return cards
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(name string) error {
	return ioutil.WriteFile(name, []byte(d.toString()), 0666)
}

func readDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		println(err)
		return nil
	}
	return deck(strings.Split(string(bs), ","))
}

func 