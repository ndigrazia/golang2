package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"math/rand"
	"time"
)

type deck []string

//Receiver function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
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

func (d deck) saveToFile(filename string) error {
 return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}


func newDeckFromFile(filename string) deck {
	var cards deck

	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		cards = newDeck()
	} else {
		cards = deck(strings.Split(string(bs), ","))
	}
   
	return cards
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}