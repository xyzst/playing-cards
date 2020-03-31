package main

func main() {
	cards := create()

	cards.print()

	cards.shuffle()

	cards.print()
}
