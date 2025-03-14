package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

//var name string = "myfile.txt"

// Execute the following commands:
// go build main.go 
//./main myfile.txt 

func main() {
	//go run main.go myfile.txt
	file, err := os.Open(os.Args[1]) // For read access.
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println(os.Args)

	io.Copy(os.Stdout, file)
}

