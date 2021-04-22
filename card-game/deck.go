package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
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

func (deck d) shuffle() {
	// Normal rand.Intn() always returns same set of integer for every execution

	// using source we can overcome the above scenario by passing the seed
	// https://gobyexample.com/random-numbers
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source);
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newnewPosition] = d[newPosition], d[i]
	}
}
