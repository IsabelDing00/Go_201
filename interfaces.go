/*
Interface
! Notice: Go does not support overloading, but Java, python, and C# will
*/

package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishbot struct{}

func main() {
	// init english bot and spanishbot
	eb := englishBot{}
	sb := spanishbot{}

	printGreeting(eb)
	printGreeting(sb)
}

// func printGreeting(eb englishBot){
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishbot){
// 	fmt.Println(sb.getGreeting())
// }

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// Assume there we have very custom logic for generating an english greeting here
	return "Hi there!"
}

func (sb spanishbot) getGreeting() string {
	return "Hola!"
}
