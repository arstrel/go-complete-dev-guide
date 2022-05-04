package main

import (
	"fmt"
)

func main() {
	s := [10]int{}

	for i := 0; i < 10; i++ {
		s[i] = i + 1
	}

	for _, val := range s {
		isEven := val%2 == 0

		if isEven {
			fmt.Printf("%v is Even\n", val)
		} else {
			fmt.Printf("%v is Odd\n", val)
		}
	}
}
