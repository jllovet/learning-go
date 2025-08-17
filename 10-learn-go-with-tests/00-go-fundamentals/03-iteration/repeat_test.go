package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat string 4 times", func(t *testing.T) {
		got := Repeat("a", 4)
		want := "aaaa"
		if got != want {
			t.Errorf("want %s but got %s", want, got)
		}
	})
	t.Run("repeating empty string returns empty string", func(t *testing.T) {
		got := Repeat("", 4)
		want := ""
		if got != want {
			t.Errorf("want %s but got %s", want, got)
		}
	})
	t.Run("repeating string 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		if got != want {
			t.Errorf("want %s but got %s", want, got)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 4)
	}
}

func BenchmarkStdLibRepeat(b *testing.B) {
	for b.Loop() {
		StdLibRepeat("a", 4)
	}
}
