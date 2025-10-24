package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gregor-pifko/learn-go-with-tests/di"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	di.Greet(w, "world")
}

func main() {
	buffer := bytes.Buffer{}
	di.Greet(&buffer, "Jodie")
	fmt.Println(buffer.String())

	di.Greet(os.Stdout, "Elodie")
	fmt.Println()

	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
