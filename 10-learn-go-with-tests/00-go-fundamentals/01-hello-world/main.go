package main

import "fmt"

const spanish = "Spanish"
const enGreetingPrefix = "Hello, "
const esGreetingPrefix = "Hola, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	switch language {
	case spanish:
		return esGreetingPrefix + name
	default:
		return enGreetingPrefix + name
	}
}

func main() {
	fmt.Println(Hello("Jonathan", ""))
}
