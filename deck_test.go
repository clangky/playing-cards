package main

import "testing"

func TestDealDoesNotExceedDeck(t *testing.T) {
	d := newDeck()
	remainder, hand := d.deal(60)
	if len(hand) != 52 {
		t.Errorf("expected hand size 52, got %d", len(hand))
	}
	if len(remainder) != 0 {
		t.Errorf("expected remainder 0, got %d", len(remainder))
	}
}

func TestDealReducesDeck(t *testing.T) {
	d := newDeck()
	remainder, hand := d.deal(5)
	if len(hand) != 5 {
		t.Errorf("expected hand size 5, got %d", len(hand))
	}
	if len(remainder) != 52-5 {
		t.Errorf("expected remainder %d, got %d", 52-5, len(remainder))
	}
}
