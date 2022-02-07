package main

import "fmt"

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	greet(eb, "English: ")
	greet(sb, "Spanish: ")

}

type bot interface {
	getGreeting(string) string
}

func greet(b bot, s string) {
	fmt.Println(b.getGreeting(s))
}

type englishBot struct{}

func (englishBot) getGreeting(s string) string {
	return s + "Hi!"
}

type spanishBot struct{}

func (spanishBot) getGreeting(s string) string {
	return s + "Hola!"
}
