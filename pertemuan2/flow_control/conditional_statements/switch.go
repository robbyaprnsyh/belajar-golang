package main

import "fmt"

func main() {
	language := "golang"

	switch language {
	case "Python":
		fmt.Println("You are learning Python!")

	case "Go", "golang":
		fmt.Println("Good, go for Go!")

	default:
		fmt.Println("Any other programing language is a good start!")
	}
}