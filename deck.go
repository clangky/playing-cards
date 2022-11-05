package main

import (
	"math/rand"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	suits := []string{"S", "D", "H", "C"}
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+suit)
		}
	}
	return cards
}

func (d deck) deal(handSize int) (deck, deck) {
	var hand deck
	for i := handSize; i > 0; i-- {
		hand = append(hand, d[0]) // one from the top
		d = d[1:]
	}
	return d, hand
}

func (d deck) shuffle() deck {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
	return d
}

