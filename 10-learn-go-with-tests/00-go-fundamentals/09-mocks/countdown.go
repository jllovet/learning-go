package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(w io.Writer, sleeper Sleeper) {
	const countdownStart = 3
	const countdownEnd = 1
	const finalWord = "Go!"
	for i := countdownStart; i >= countdownEnd; i-- {
		fmt.Fprintf(w, "%d\n", i)
		sleeper.Sleep()
	}
	fmt.Fprint(w, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{
		duration: 1 * time.Second,
		sleep:    time.Sleep,
	}
	Countdown(os.Stdout, sleeper)
}
