package smi

import "testing"

func assertCorrectMessage(t testing.TB, got, want any) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
