package main 

// import "fmt"

func main() {
	// cards := newDeck()

	// cards.saveToFile("my_cards")

	cards := newDeckFromFile("my_card")
	cards.print()
}
