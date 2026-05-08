package poker_test

import (
	"io"
	"testing"

	"github.com/gregor-pifko/learn-go-with-tests/poker"
)

func TestTape_Write(t *testing.T) {
	file := createTempFile(t, "12345")

	tape := poker.Tape{file}

	tape.Write([]byte("abc"))

	file.Seek(0, io.SeekStart)
	newFileContents, _ := io.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("error while writing to file, got %q; want %q", got, want)
	}
}
