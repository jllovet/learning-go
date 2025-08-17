package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	repeated := Repeat("apple", 4)
	fmt.Println(repeated)
	// Output: appleappleappleapple
}

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
	t.Run("repeating long string", func(t *testing.T) {
		got := Repeat("apple", 4)
		want := "appleappleappleapple"
		if got != want {
			t.Errorf("want %s but got %s", want, got)
		}
	})
}

func TestBuilderRepeat(t *testing.T) {
	t.Run("repeat string 4 times", func(t *testing.T) {
		got := BuilderRepeat("a", 4)
		want := "aaaa"
		if got != want {
			t.Errorf("want %s but got %s", want, got)
		}
	})
	t.Run("repeating empty string returns empty string", func(t *testing.T) {
		got := BuilderRepeat("", 4)
		want := ""
		if got != want {
			t.Errorf("want %s but got %s", want, got)
		}
	})
	t.Run("repeating string 5 times", func(t *testing.T) {
		got := BuilderRepeat("a", 5)
		want := "aaaaa"
		if got != want {
			t.Errorf("want %s but got %s", want, got)
		}
	})
	t.Run("repeating long string", func(t *testing.T) {
		got := BuilderRepeat("apple", 4)
		want := "appleappleappleapple"
		if got != want {
			t.Errorf("want %s but got %s", want, got)
		}
	})
}

func TestStdLibRepeat(t *testing.T) {
	t.Run("repeat string 4 times", func(t *testing.T) {
		got := StdLibRepeat("a", 4)
		want := "aaaa"
		if got != want {
			t.Errorf("want %s but got %s", want, got)
		}
	})
	t.Run("repeating empty string returns empty string", func(t *testing.T) {
		got := StdLibRepeat("", 4)
		want := ""
		if got != want {
			t.Errorf("want %s but got %s", want, got)
		}
	})
	t.Run("repeating string 5 times", func(t *testing.T) {
		got := StdLibRepeat("a", 5)
		want := "aaaaa"
		if got != want {
			t.Errorf("want %s but got %s", want, got)
		}
	})
	t.Run("repeating long string", func(t *testing.T) {
		got := StdLibRepeat("apple", 4)
		want := "appleappleappleapple"
		if got != want {
			t.Errorf("want %s but got %s", want, got)
		}
	})
}

func BenchmarkRepeatShortString(b *testing.B) {
	for b.Loop() {
		Repeat("a", 4)
	}
}
func BenchmarkRepeatLongString(b *testing.B) {
	for b.Loop() {
		Repeat("apple", 4)
	}
}
func BenchmarkBuilderRepeatShortString(b *testing.B) {
	for b.Loop() {
		BuilderRepeat("a", 4)
	}
}
func BenchmarkBuilderRepeatLongString(b *testing.B) {
	for b.Loop() {
		BuilderRepeat("apple", 4)
	}
}

func BenchmarkStdLibRepeatShortString(b *testing.B) {
	for b.Loop() {
		StdLibRepeat("a", 4)
	}
}
func BenchmarkStdLibRepeatLongString(b *testing.B) {
	for b.Loop() {
		StdLibRepeat("apple", 4)
	}
}
