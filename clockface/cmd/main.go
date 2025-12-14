package main

import (
	"os"
	"time"

	"github.com/gregor-pifko/learn-go-with-tests/clockface/svg"
)

func main() {
	svg.Write(os.Stdout, time.Now())
}
