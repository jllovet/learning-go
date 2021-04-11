package main

import "fmt"

func main() {
	car, cost := "Audi", 200 // declaration
	// car, cost := "Mazda", 300 // results in compile-time error because there is no new variable declared
	car, cost = "Mazda", 300 // assignment
	fmt.Println(car)
	fmt.Println(cost)
}
