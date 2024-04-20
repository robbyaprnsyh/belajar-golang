package main

import (
	"fmt"
	"time"
)

func main() {
	hour := time.Now().Hour()

	switch {
	case hour < 12:
		fmt.Println("Good Morning")
	case hour < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Evening!")
	}
}