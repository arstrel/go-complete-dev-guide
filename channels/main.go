package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://stackoverflow.com",
		"http://facebook.com",
		"http://artemio.tech",
	}

	for _, val := range links {
		wg.Add(1)
		go checkLink(val)
	}

	wg.Wait()
}

func checkLink(link string) {
	defer wg.Done()
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
}
