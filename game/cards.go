package game

import "fmt"

// func main() {
// 	d := newDeck()
// 	var size int = len(d)
// 	var handOne deck
// 	var deckSize int
// 	var lastDeck deck
// 	handOne, lastDeck, deckSize = d.fiveCardDeal()
// 	var handTwo []string
// 	handTwo, lastDeck, deckSize = d.fiveCardDeal()
// 	fmt.Println(size)
// 	fmt.Println(handOne)
// 	fmt.Println(handTwo)
// 	fmt.Println(deckSize)
// 	fmt.Println(lastDeck)
// }

func game() {
	d := newDeck()
	var size int = len(d)
    fmt.Println(size)
	d = d.shuffle()
	stack, handOne := d.deal(7)
	size = len(stack)
	fmt.Println(handOne)
	fmt.Println(stack)
    fmt.Println(size)
}

