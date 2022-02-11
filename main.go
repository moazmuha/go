package main

import "fmt"


type myInt int
func main() {
	fmt.Println(`Hello World!`)
	// if x := test(); x {
	// 	fmt.Println(x)
	// }

	switch {
		case true:
			fmt.Println(`Hello World`)
		case false:
			if false {fmt.Println(`Hello World1`)}
		case false:
			fmt.Println(`Hello World2`)
		default:
			fmt.Println(`Default`)

	}
	c := make(chan string)

	go test(c)

	// msg :=  <- c
	
	fmt.Println(<- c)

	// m := map[string]map[string]string{

	// }

	fmt.Println("hi", 5)
}


func test(c chan string) bool {
	c <- "hi"
	return true

}