package main

import "fmt"

func main() {
	names := []string{"Jhon", "Paul", "George", "Ringo"}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
	
}