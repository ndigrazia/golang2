package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(resp)

	print(resp.Body)
}

func print(r io.Reader) {
	bs := make([]byte, 99999)
	r.Read(bs)
	fmt.Println(string(bs))
}

