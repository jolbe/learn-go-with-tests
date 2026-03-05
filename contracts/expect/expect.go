// ABOUTME: Test assertion helpers for contracts tests.
// ABOUTME: Provides NoErr, Err, and Equal assertion functions.
package expect

import (
	"errors"
	"testing"
)

func NoErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func Err(t *testing.T, got, target error) {
	t.Helper()
	if got == nil {
		t.Fatal("expected an error, got nil")
	}
	if !errors.Is(got, target) {
		t.Fatalf("expected error %v, got %v", target, got)
	}
}

func Equal[T comparable](t *testing.T, expected, got T) {
	t.Helper()
	if expected != got {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}
