package main

import "fmt"

func main() {
	mySlice := []string{"Hi", "There", "how", "are", "you"}
	fmt.Println(mySlice)

	update(mySlice)
	fmt.Println(mySlice)
}

func update(s []string) {
	s[0] = "Hello"
}