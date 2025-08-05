package main

import "fmt"

const enGreetingPrefix = "Hello, "

func Hello(name string) string {
	switch name {
	case "":
		return enGreetingPrefix + "World"
	default:
		return enGreetingPrefix + name
	}
}

func main() {
	fmt.Println(Hello("Jonathan"))
}
