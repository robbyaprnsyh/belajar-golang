package main

import "fmt"

func getName() (firstName, lastName string) {
	return "Robby", "Apriansyah"
}

func main() {
	x, y := getName()
	fmt.Println(x, y)
}
