package main

import "fmt"

func main() {
	c := make(chan int)

	go func(){
		c <- 17
		close(c)
	}()

	func(){
		fmt.Println(<- c)
	}()
	
	i, ok := <-c

	fmt.Println(i, ok)
	fmt.Println("done")
}