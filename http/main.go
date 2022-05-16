package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://www.google.com")

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	// Approach 1
	// bs := make([]byte, 999999)
	// resp.Body.Read(bs)

	// Approach 2
	// body, err := io.ReadAll(resp.Body)
	// resp.Body.Close()
	// if err != nil {
	// 	fmt.Println("Error")
	// }

	// Approach 3
	io.Copy(os.Stdout, resp.Body)

	// Approach 4: custom writer
	lw := logWriter{}
	io.Copy(lw, resp.Body)

	// fmt.Printf("%s", body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
