package main

import (
	"fmt"
	"github.com/otiai10/gosseract/v2"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()

	err := client.SetImage("/Users/bigbo/Desktop/1111.png")
	fmt.Println(err)

	text, _ := client.Text()
	fmt.Println(text)
	// Hello, World!
}