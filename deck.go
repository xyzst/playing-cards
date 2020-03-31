package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// Extend 'slice of string' as new type called 'deck'
type deck []string

// create will generate a full 52-card deck
func create() deck {
	cards := deck{}

	suits := []string{"♦️", "♥️", "♠️️", "♣️"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// print will log the current contents to stdout
func (d deck) print() { // known as a "receiver" on a function, similar to "this"/"self" in other languages. "d" is a reference (not a copy)
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// toString will return a string representation of the current deck, each string separated by a comma (,)
func (d deck) toString() string {
	return strings.Join(d, ",")
}

// serialize will save the current deck to disk given the fileName
func (d deck) serialize(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

// deal will create a single hand from a specified deck
func deal(d deck, size int) (deck, deck) {
	return d[:size], d[size:]
}

// deserialize will read from disk given the fileName and return a hydrated deck instance
func deserialize(fileName string) (deck, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("error: ", err)
		return nil, err
	}
	return strings.Split(string(bytes), ","), nil
}

// shuffle will rearrange the ordering of the deck. Performs side effects on the current deck object
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) // Is there a better seed/source for RNG?
	r := rand.New(source)
	for c := range d {
		i := r.Intn(len(d) - 1)
		d[c], d[i] = d[i], d[c]
	}
}
