package main

import (
	"math/rand"
	"time"
)

type card struct {
	suit  string
	value string
}

type deck []card

func newDeck() deck {
	cards := deck{}

	suits := []string{"S", "D", "H", "C"}
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, card{suit, value})
			// cards = append(cards, value+suit)
		}
	}
	return cards
}

func (d *deck) deal(handSize int) deck {
	if handSize > len(*d) {
		handSize = len(*d)
	}
	hand := make(deck, 0, handSize)
	for i := 0; i < handSize; i++ {
		hand = append(hand, (*d)[0]) // one from the top
		*d = (*d)[1:]
	}
	return hand
}

func (d deck) toStringSlice() []string {
	s := make([]string, len(d))
	for i := range d {
		s[i] = d[i].value + d[i].suit
	}
	return s
}

func (d deck) shuffle() deck {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
	return d
}
