package main

import "fmt"

func main() {

	numbers := map[string]int{
		"one": 1,
		"two": 2,
	}

	numbers["three"] = 3
	// delete(numbers, "one")

	fmt.Println(numbers)
	iterate(numbers)
  
}

func iterate(m map[string]int) {

	for k, v := range m {
		fmt.Println(k, v)
	}
}
