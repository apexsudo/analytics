package analytics_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/apexsudo/analytics"
)

func TestNew(t *testing.T) {
	t.Parallel()
	t.Run("returns a client", func(t *testing.T) {
		t.Parallel()
		a := analytics.New("invalid_test_key", "https://dataplane.example.com")
		assert.NotNil(t, a)
	})

	t.Run("satisfies identify call", func(t *testing.T) {
		t.Parallel()
		a := analytics.New("invalid_test_key", "https://dataplane.example.com")
		err := a.Identify("user_1", map[string]any{
			"foo":      "bar",
			"username": "alice",
		})
		assert.NoError(t, err)
	})

	t.Run("satisfies track call", func(t *testing.T) {
		t.Parallel()
		a := analytics.New("invalid_test_key", "https://dataplane.example.com")
		err := a.Track("user_1", "User logged in", map[string]any{
			"foo":      "bar",
			"username": "alice",
		})
		assert.NoError(t, err)
	})

	t.Run("can close the client", func(t *testing.T) {
		t.Parallel()
		a := analytics.New("invalid_test_key", "https://dataplane.example.com")
		assert.NoError(t, a.Close())
	})
}
