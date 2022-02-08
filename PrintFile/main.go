package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) > 1 {

		contents, err := os.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		} else {
			fmt.Println(string(contents))
		}

		fmt.Println("Done")
	}
}