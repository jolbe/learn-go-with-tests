package main_test

import (
	"testing"

	"github.com/gpifko/go-specs-greet/adapters"
	"github.com/gpifko/go-specs-greet/adapters/grcpserver"
	"github.com/gpifko/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	addr := adapters.StartDockerServer(t, "50052", "grcpserver")
	driver := grcpserver.Driver{Addr: addr}
	t.Cleanup(driver.Close)

	specifications.GreetSpecification(t, &driver)
	specifications.CurseSpecification(t, &driver)
}
