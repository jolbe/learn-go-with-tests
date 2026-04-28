package httpserver_test

import (
	"io"
	"testing"

	httpserver "github.com/gregor-pifko/learn-go-with-tests/http-server"
)

func TestTape_Write(t *testing.T) {
	file := createTempFile(t, "12345")

	tape := httpserver.Tape{file}

	tape.Write([]byte("abc"))

	file.Seek(0, io.SeekStart)
	newFileContents, _ := io.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("error while writing to file, got %q; want %q", got, want)
	}
}
