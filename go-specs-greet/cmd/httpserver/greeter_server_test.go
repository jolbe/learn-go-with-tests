package main_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gpifko/go-specs-greet/adapters"
	"github.com/gpifko/go-specs-greet/adapters/httpserver"
	"github.com/gpifko/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	addr := adapters.StartDockerServer(t, "8080/tcp", "httpserver")
	driver := httpserver.Driver{BaseURL: fmt.Sprintf("http://%s", addr), Client: &http.Client{
		Timeout: 1 * time.Second,
	}}

	specifications.GreetSpecification(t, driver)
	specifications.CurseSpecification(t, driver)
}
