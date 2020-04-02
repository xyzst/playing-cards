package main

type card struct {
	value string
	suit string
}

func (c card) toString() string {
	return c.value + " of " + c.suit
}