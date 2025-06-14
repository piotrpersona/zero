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
}
