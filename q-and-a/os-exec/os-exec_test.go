package main

import (
	"strings"
	"testing"
)

func TestGetDataIntegration(t *testing.T) {
	got := GetData(getXMLFromCommand())
	want := "TEST MESSAGE"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGetData(t *testing.T) {
	input := strings.NewReader(`
<payload>
    <message>Cats are the best animal</message>
</payload>`)

	got := GetData(input)
	want := "CATS ARE THE BEST ANIMAL"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
