package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	linkList := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range linkList {
		go checkLink(link, c)
	}

	// Note: Never access the same variable between routines (including main routine)
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, channel chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		channel <- link
		return
	}

	fmt.Println(link, "is up!")
	channel <- link
}
