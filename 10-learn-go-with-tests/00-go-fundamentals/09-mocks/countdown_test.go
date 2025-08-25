package main

import (
	"bytes"
	"testing"
)

type spySleeper struct {
	Calls int
}

func (s *spySleeper) Sleep() {
	s.Calls++
}

func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeperSpy := &spySleeper{}
	Countdown(buffer, sleeperSpy)
	got := buffer.String()
	want := `3
2
1
Go!`
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
	if sleeperSpy.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want 3 got %d", sleeperSpy.Calls)
	}
}
