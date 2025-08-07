package main

import "testing"

func TestDealDoesNotExceedDeck(t *testing.T) {
	d := newDeck()
	hand := d.deal(60)
	if len(hand) != 52 {
		t.Errorf("expected hand size 52, got %d", len(hand))
	}
	if len(d) != 0 {
		t.Errorf("expected remainder 0, got %d", len(d))
	}
}

func TestDealReducesDeck(t *testing.T) {
	d := newDeck()
	hand := d.deal(5)
	if len(hand) != 5 {
		t.Errorf("expected hand size 5, got %d", len(hand))
	}
	if len(d) != 52-5 {
		t.Errorf("expected remainder %d, got %d", 52-5, len(d))
	}
}
