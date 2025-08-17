package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const enGreetingPrefix = "Hello, "
const esGreetingPrefix = "Hola, "
const frGreetingPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	switch language {
	case spanish:
		return esGreetingPrefix + name
	case french:
		return frGreetingPrefix + name
	default:
		return enGreetingPrefix + name
	}
}

func main() {
	fmt.Println(Hello("Jonathan", ""))
}
