package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(w io.Writer, name string) {
	// message := fmt.Sprintf("Hello, %s\n", name)
	// w.Write([]byte(message))
	fmt.Fprintf(w, "Hello, %s", name)
}

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(GreetHandler)))
}
