package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(w io.Writer) {
	fmt.Fprintf(w, "3\n")
	fmt.Fprintf(w, "2\n")
	fmt.Fprintf(w, "1\n")
	fmt.Fprintf(w, "Go!")
}

func main() {
	Countdown(os.Stdout)
}
