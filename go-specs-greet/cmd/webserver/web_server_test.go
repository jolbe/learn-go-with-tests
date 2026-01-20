package main_test

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/gpifko/go-specs-greet/adapters"
	"github.com/gpifko/go-specs-greet/adapters/webserver"
	"github.com/gpifko/go-specs-greet/specifications"
)

func TestGreeterWeb(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	addr := adapters.StartDockerServer(t, "8081", "webserver")
	driver, cleanup := webserver.NewDriver(fmt.Sprintf("http://%s", addr))

	t.Cleanup(func() {
		assert.NoError(t, cleanup())
	})

	specifications.GreetSpecification(t, driver)
	specifications.CurseSpecification(t, driver)
}
