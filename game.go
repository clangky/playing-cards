package main

import (
	"math/rand"
	"time"
)

const (
	startingChips = 20000
	smallBlind    = 100
	bigBlind      = 200
)

type Player struct {
	Name   string
	Hand   deck
	Chips  int
	InHand bool
}

type Game struct {
	deck    deck
	board   deck
	players [2]*Player
	pot     int
	stage   int // 0 preflop,1 flop,2 turn,3 river,4 showdown
}

type State struct {
	PlayerChips int      `json:"playerChips"`
	NPCChips    int      `json:"npcChips"`
	Pot         int      `json:"pot"`
	Board       []string `json:"board"`
	PlayerHand  []string `json:"playerHand"`
	NPCHand     []string `json:"npcHand"`
	Message     string   `json:"message"`
}

func NewGame() *Game {
	d := newDeck().shuffle()
	g := &Game{
		deck:  d,
		board: deck{},
		players: [2]*Player{
			{Name: "You", Chips: startingChips, InHand: true},
			{Name: "Dali NPC", Chips: startingChips, InHand: true},
		},
		stage: 0,
	}
	g.deal()
	return g
}

func (g *Game) deal() {
	h := g.deck.deal(2)
	g.players[0].Hand = h
	h = g.deck.deal(2)
	g.players[1].Hand = h
	g.players[0].Chips -= smallBlind
	g.players[1].Chips -= bigBlind
	g.pot = smallBlind + bigBlind
}

func (g *Game) nextStage() {
	switch g.stage {
	case 0:
		h := g.deck.deal(3)
		g.board = append(g.board, h...)
	case 1, 2:
		h := g.deck.deal(1)
		g.board = append(g.board, h...)
	case 3:
		winner := g.winner()
		g.players[winner].Chips += g.pot
		g.stage = 4
		return
	}
	g.stage++
}

var rank = map[string]int{
	"A":  14,
	"K":  13,
	"Q":  12,
	"J":  11,
	"10": 10,
	"9":  9,
	"8":  8,
	"7":  7,
	"6":  6,
	"5":  5,
	"4":  4,
	"3":  3,
	"2":  2,
}

func highCard(h deck, b deck) int {
	max := 0
	cards := append(h, b...)
	for _, c := range cards {
		if r := rank[c.value]; r > max {
			max = r
		}
	}
	return max
}

func (g *Game) winner() int {
	hc0 := highCard(g.players[0].Hand, g.board)
	hc1 := highCard(g.players[1].Hand, g.board)
	if hc0 >= hc1 {
		return 0
	}
	return 1
}

func (g *Game) PlayerFold() string {
	g.players[0].InHand = false
	g.players[1].Chips += g.pot
	g.stage = 4
	return "You folded."
}

func (g *Game) PlayerCall() string {
	if g.stage == 0 {
		need := bigBlind - smallBlind
		g.players[0].Chips -= need
		g.pot += need
	}
	return "You call."
}

func (g *Game) NPCAct() string {
	rand.Seed(time.Now().UnixNano())
	quips := []string{
		"Is that the best you've got?",
		"I've seen snails think faster.",
		"Your play is melting like a clock.",
		"Even my dreams are tougher opponents.",
	}
	if g.stage < 4 {
		g.nextStage()
	}
	return quips[rand.Intn(len(quips))]
}

func (g *Game) State(revealNPC bool, msg string) State {
	npcHand := []string{}
	if revealNPC {
		npcHand = g.players[1].Hand.toStringSlice()
	}
	return State{
		PlayerChips: g.players[0].Chips,
		NPCChips:    g.players[1].Chips,
		Pot:         g.pot,
		Board:       g.board.toStringSlice(),
		PlayerHand:  g.players[0].Hand.toStringSlice(),
		NPCHand:     npcHand,
		Message:     msg,
	}
}

func (g *Game) HandOver() bool {
	return g.stage == 4
}
