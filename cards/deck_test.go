package main

import "testing"

func testNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 2000 {
		t.Error("Error")
	}
}
