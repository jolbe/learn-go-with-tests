package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expect %q; got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 500)
	}
}

func ExampleRepeat() {
	repeated := Repeat("A", 5)
	fmt.Println(repeated)
	// Output: AAAAA
}
