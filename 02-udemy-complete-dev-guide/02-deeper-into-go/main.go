package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// go is statically typed!
	card := "Ace of Spades"  // Initialization
	card = "Queen of Hearts" // Reassignment (note the difference in operator)
	fmt.Println(card)
}
