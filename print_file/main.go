package main

import (
	"fmt"
	"io"
	"os"
)

// Approach 2
func main() {
	fileName := os.Args[1]
	fmt.Println(fileName)

	f, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	content, err := io.ReadAll(f)

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	fmt.Println(string(content))
}

// Approach 1
// func main() {
// 	fileName := os.Args[1]
// 	fmt.Println(fileName)

// 	content, err := ioutil.ReadFile(fileName)

// 	if err != nil {
// 		fmt.Println("Error: ", err)
// 		os.Exit(1)
// 	}

// 	fmt.Println(string(content))
// }
