package main

//import "fmt"

func main() {
	//Declare & assign
	//var card string = "ace of spades"
	//card := "ace of spades"
	//card = "five of diamonds"

	//Functions
	//card := newCard()
	
	//card = newCard2("two of spades")

	//Array
	/*cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
		
	fmt.Println(cards)

	*/

	/*Types
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()*/

	/*
	cards := newDeck()
	cards.print()
	*/

	/*cards := newDeck()
	hand, remainingCards := cards.deal(3)

	fmt.Println("Hand:")
	hand.print()

	fmt.Println("Remaining Cards:")
	remainingCards.print()*/

	/*cards := newDeck()
	fmt.Println(cards.toString())*/

	/*cards := newDeck()
	cards.saveToFile("my_cards")*/

	/*cards := newDeckFromFile("my_cards")
	cards.print()*/

	cards := newDeck()
	cards.shuffle()
	cards.print()

}

func newCard() string {
	return "five of diamonds"
}

func newCard2(card string) string {
	return card
}
