package main

import "fmt"

/*
//Without interfaces
type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	//Custom logic
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	//Custom logic
	return "Hola!"
}

func printGreeting(b englishBot) {
	fmt.Println(b.getGreeting())
}

//NO identical name - No overloading
func printGreeting2(b spanishBot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting2(sb)
}
*/

//With interfaces
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	//Custom logic
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	//Custom logic
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}
