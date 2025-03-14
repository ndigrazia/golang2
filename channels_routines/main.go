package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	
	//msg := <-c
	//fmt.Println(msg)

	//fmt.Println(<-c)

	/*for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}*/

	/*for {
		go checkLink(<-c, c)
	}*/

	/*for l := range c {
		checkLink(l, c)
	}*/

	/*
	//No reference to a outer variable l. It's a wrong use practice.
	for l := range c {
		go func() {
			time.Sleep(time.Second)
			checkLink(l, c)
		}()
	}*/	

	for l := range c {
		go func(link string) {
			time.Sleep(time.Second)
			checkLink(link, c)
		}(l)

		//go delay(l, c)
	}	
}

func delay(l string, c chan string) {	
	time.Sleep(time.Second * 5)
	checkLink(l, c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}