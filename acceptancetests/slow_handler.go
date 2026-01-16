// Package acceptancetests is a simple HTTP handler
package acceptancetests

import (
	"fmt"
	"net/http"
	"time"
)

func SlowHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("slow handling request...")
	time.Sleep(2 * time.Second)
	w.Write([]byte("Hello world let's see if this actually works"))
}
