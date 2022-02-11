package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	sites := []string{"http://google.com", "http://github.com/moazmuha", "http://linkedin.com/moazmuha", "http://stackoverflow.com"}
	c := make(chan string)

	for _, site := range sites {
		go statusChecker(site, c)
	}

	for msg := range c {
		go func(msg string) {
			time.Sleep(2 * time.Second)
			statusChecker(msg, c)
		}(msg)
	}
}

func statusChecker(site string, c chan string) {
	_, err := http.Get(site)

	if err == nil {
		fmt.Println("Seems like", site, "is up")
	} else {
		fmt.Println("Seems like", site, "is down")
		fmt.Println("Error:", err)
	}

	c <- site

}
