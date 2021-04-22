package main

func main() {
	cards := readDeckFromFile("deck-1.txt")
	println(cards.toString())
}
