package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	myctx "github.com/gregor-pifko/learn-go-with-tests/context"
)

type MyStore struct {
	data string
}

func (s *MyStore) Fetch(ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		fmt.Println("Fetch got cancelled")
		return "", ctx.Err()
	case <-time.After(3 * time.Second):
		return "got: " + s.data, nil
	}
}

func main() {
	store := &MyStore{"This is sample server"}
	ctx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(2*time.Second, cancel)

	go http.ListenAndServe(":8080", myctx.Server(store))

	// failed attempt
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/", nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
		// return
	}

	// happy path
	req, _ = http.NewRequestWithContext(context.Background(), http.MethodGet, "http://localhost:8080/", nil)
	client = &http.Client{}
	resp, _ = client.Do(req)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
