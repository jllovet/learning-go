package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(w io.Writer) {
	const countdownStart = 3
	const countdownEnd = 1
	const finalWord = "Go!"
	for i := countdownStart; i >= countdownEnd; i-- {
		fmt.Fprintf(w, "%d\n", i)
	}
	fmt.Fprint(w, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
