package main

import "fmt"

func logging() {
	fmt.Println("Finish running the program")
}

func runApplication() {
	fmt.Println("Run application")
}

func main() {
	defer logging()
	runApplication()
}
