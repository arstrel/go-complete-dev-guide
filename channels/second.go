package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Example with waitGroup

var wg_one sync.WaitGroup

func mainOne() {
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://stackoverflow.com",
		"http://facebook.com",
		"http://artemio.tech",
	}

	for _, val := range links {
		wg_one.Add(1)
		go checkLinkOne(val)
	}

	wg_one.Wait()
}

func checkLinkOne(link string) {
	defer wg_one.Done()
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
}
