package adapters

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/alecthomas/assert/v2"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func StartDockerServer(t testing.TB, exposedPort, binToBuild string) string {
	ctx := context.Background()
	t.Helper()

	// TODO: this is hardcoded and the paths are relative to where you run go test from
	// which is the package directory ./cmd/httpserver
	// Possible resolutions:
	// - pass contextPath as parameter too
	// - use an absolute path based on runtime.Caller
	// - root module detection
	df := testcontainers.FromDockerfile{
		Context:    "../../.",
		Dockerfile: "Dockerfile",
		BuildArgs: map[string]*string{
			"bin_to_build": &binToBuild,
		},
		BuildLogWriter: os.Stderr,
	}
	container, err := testcontainers.Run(
		ctx,
		"",
		testcontainers.WithDockerfile(df),
		testcontainers.WithExposedPorts(exposedPort),
		testcontainers.WithWaitStrategy(wait.ForListeningPort(nat.Port(exposedPort))),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		assert.NoError(t, container.Terminate(cleanupCtx))
	})

	host, err := container.Host(ctx)
	assert.NoError(t, err)
	port, err := container.MappedPort(ctx, nat.Port(exposedPort))
	assert.NoError(t, err)

	return host + ":" + port.Port()
}
