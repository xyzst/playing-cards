package main

import "testing"

// go has a built in testing framework, each test must be formatted as such: "TestSomeDescriptionHere(t *testing.T)"
func TestNewDeckShouldReturnExactly52Cards(t *testing.T) {
	deck := create()

	if len(deck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(deck))
	}
}
