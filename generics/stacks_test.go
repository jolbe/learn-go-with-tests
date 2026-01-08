package generics_test

import (
	"testing"

	"github.com/gregor-pifko/learn-go-with-tests/generics"
)

func TestStackOfStrings(t *testing.T) {
	stack := generics.StackOfStrings{}

	// Test empty pop
	_, ok := stack.Pop()
	if ok {
		t.Error("expected false when popping empty stack")
	}

	// Test push/pop
	stack.Push("hello")
	stack.Push("world")

	got, ok := stack.Pop()
	if !ok {
		t.Error("expected true when popping non-empty stack")
	}
	if got != "world" {
		t.Errorf("got %q; want %q", got, "world")
	}

	got, ok = stack.Pop()
	if !ok {
		t.Error("expected true on second pop")
	}
	if got != "hello" {
		t.Errorf("got %q; want %q", got, "hello")
	}
}
