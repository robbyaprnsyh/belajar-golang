package main

import "fmt"

func main() {
	numbers := []int{2, 3, 4, 5}
	fmt.Println(numbers)

	nums := make([]int, 2)
	fmt.Printf("%#v\n", nums)

	friends := [3]string{"John", "Paul", "George"}
	fmt.Printf("%#v\n", friends)

	myBestFriends := friends[0:1]
	fmt.Println("My best friend is ", myBestFriends)
	fmt.Printf("%#v\v", myBestFriends)

	for index, value := range friends {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}

}