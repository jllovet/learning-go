package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	want := Add(2, 2)
	got := 4

	if want != got {
		t.Errorf("wanted '%d' but got '%d'", want, got)
	}
}

func ExampleAdd() {
	sum := Add(2, 2)
	fmt.Println(sum)
	// Output: 4
}
