package main

import "fmt"

func main() {
	//Initializing a map / Alternative 1
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	
	printMap(colors)


	/*
	//Initializing a map / Alternative 2
	colors := make(map[int]string)

	colors[10] = "#ffffff"

	delete(colors, 10)

	fmt.Println(colors)*/

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

