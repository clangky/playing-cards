package game

// import "fmt"
import (
	"math/rand"
	"time"
)

// Create a new type 'Deck'
// which is a slice of type string

type deck []string

func newDeck() deck {
	cards := deck{}

	suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	nums := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	// nums := []string{"Ace", "Two", "Three"}

	for _, suit := range suits {
		for _, value := range nums {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) deal(handSize int) (deck, deck) {
	var hand deck
	for i := handSize; i > 0; i-- {
		hand = append(hand, d[0]) // one from the top of
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


func (d deck) randDeal(rando int) (deck, string) {
	nextDeck := append(d[:rando], d[rando+1:]...)
	return nextDeck, d[rando]
}

// func (d deck) randomCard() int {
// 	rand.Seed(time.Now().UnixNano())
// 	return rand.Intn(len(d))
// }

// func (d deck)fiveCardDeal() (deck, deck, int) {
// 	var hand deck
// 	var remaDeck deck
// 	var card string
// 	for i := 0; i < 5; i++ {
// 		var rnd int = d.randomCard()
// 		remaDeck, card = d.randDeal(rnd)
// 		hand = append(hand, card)
// 	}
// 	var deckSize int = len(remaDeck)
// 	return hand, remaDeck, deckSize
// }

