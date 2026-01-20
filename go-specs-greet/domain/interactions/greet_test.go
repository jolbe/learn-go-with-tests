package interactions_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/gpifko/go-specs-greet/domain/interactions"
	"github.com/gpifko/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(interactions.Greet),
	)

	t.Run("default name to \"World\" if it's an empty string", func(t *testing.T) {
		assert.Equal(t, "Hello, World", interactions.Greet(""))
	})
}
