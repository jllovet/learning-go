package main

import "fmt"

const (
	spanish          = "Spanish"
	french           = "French"
	enGreetingPrefix = "Hello, "
	esGreetingPrefix = "Hola, "
	frGreetingPrefix = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frGreetingPrefix
	case spanish:
		prefix = esGreetingPrefix
	default:
		prefix = enGreetingPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("Jonathan", ""))
}
