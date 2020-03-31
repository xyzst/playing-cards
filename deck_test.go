package main

import (
	"os"
	"testing"
)

// go has a built in testing framework, each test must be formatted as such: "TestSomeDescriptionHere(t *testing.T)"
func TestNewDeckShouldReturnExactly52Cards(t *testing.T) {
	deck := create()

	if len(deck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(deck))
	}
}

func TestSerializationToFileAndDeserializationToDeck(t *testing.T) {
	// Remove any remnants from previous tests or test run
	os.Remove("_myDeck")

	deck := create()

	err := deck.serialize("_myDeck")

	if err != nil {
		t.Errorf("Unable to serialize deck due to %v", err)
	}

	saved, err := deserialize("_myDeck")

	if err != nil {
		t.Errorf("Unable to deserialize byte slice to deck object due to %v", err)
	}

	if len(saved) != 52 {
		t.Errorf("Expected 52 cards to be deserialized, but got %v", len(saved))
	}

	// Clean up
	os.Remove("_myDeck")
}
