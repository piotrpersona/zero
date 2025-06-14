package optional_test

import (
	"testing"

	"github.com/piotrpersona/zero/optional"
	"github.com/stretchr/testify/require"
)

func Test_Optional(t *testing.T) {
	t.Parallel()

	t.Run("Test some", func(t *testing.T) {
		t.Parallel()

		opt := optional.Some("car")
		val, err := opt.Get()
		require.NoError(t, err)
		require.Equal(t, "car", val)
	})
	t.Run("Test none", func(t *testing.T) {
		t.Parallel()

		opt := optional.None[string]()
		val, err := opt.Get()
		require.ErrorIs(t, err, optional.NoneErr)
		require.Equal(t, "", val)
	})
	t.Run("Test default", func(t *testing.T) {
		t.Parallel()

		opt := optional.None[string]()
		val := opt.Default("boat")
		require.Equal(t, "boat", val)
	})
	t.Run("Test From", func(t *testing.T) {
		t.Parallel()

		a := 32
		opt := optional.From(&a)
		val, err := opt.Get()
		require.NoError(t, err)
		require.Equal(t, a, val)
	})
	t.Run("Test From nil", func(t *testing.T) {
		t.Parallel()

		var p *int
		opt := optional.From(p)
		_, err := opt.Get()
		require.ErrorIs(t, err, optional.NoneErr)
	})
	t.Run("Test FromDefault", func(t *testing.T) {
		t.Parallel()

		var p *int
		opt := optional.FromDefault(p, 32)
		val, err := opt.Get()
		require.NoError(t, err)
		require.Equal(t, 32, val)
	})
}
