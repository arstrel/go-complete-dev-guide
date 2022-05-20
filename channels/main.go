package main

import (
	"fmt"
	"net/http"
	"time"
)

// Site status checker

// Example with a channel

func main() {
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://stackoverflow.com",
		"http://facebook.com",
		"http://artemio.tech",
	}

	c := make(chan string)

	for _, val := range links {
		go checkLink(val, c)
	}

	// // infinite loop. Approach 1
	// for {
	// 	// this will block for loop from continuing
	// 	// so it is "waiting" for messages coming from the changel
	// 	go checkLink(<-c, c)

	// }

	// loop as long as values coming from the channel. Approach 2
	for l := range c {
		// this will block for loop from continuing
		// so it is "waiting" for messages coming from the changel

		// "function literal" = lambda or anonymous function
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)

	}

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
