package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const sleep = "sleep"
const write = "write"

type spyCountdownOperations struct {
	Calls []string
}

func (s *spyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *spyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type spyTime struct {
	durationSlept time.Duration
}

func (s *spyTime) setDurationSlept(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	t.Run("", func(t *testing.T) {
		sleepTime := 5 * time.Second
		spy := &spyTime{}
		sleeper := ConfigurableSleeper{sleepTime, spy.setDurationSlept}
		sleeper.Sleep()
		if spy.durationSlept != sleepTime {
			t.Errorf("slept for %v but should have slept for %v", spy.durationSlept, sleepTime)
		}
	})
}

func TestCountDown(t *testing.T) {
	t.Run("print 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &spyCountdownOperations{})
		got := buffer.String()
		want := `3
2
1
Go!`
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("print and then sleep before every subsequent print", func(t *testing.T) {
		spySleepPrinter := &spyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v but got %v", want, spySleepPrinter.Calls)
		}
	})

}
