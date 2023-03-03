package main

import "fmt"

func main() {
    var card string = "Five of Diamonds"
    card = "Ace of Spades"
    // card := newCard()

    fmt.Println(card)
}

func newCard() string {
   return "Five of Diamonds" 
}
