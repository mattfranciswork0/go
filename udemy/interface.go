package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string //any type using getGreeting() can use the bot method
}

func runInterfaceExample() {
	sb := spanishBot{}
	eb := englishBot{}
	eb.getGreeting()
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
func (eb englishBot) getGreeting() string {
	return "Hi"
}

func (sb spanishBot) getGreeting() string {
	return "Hi"
}
