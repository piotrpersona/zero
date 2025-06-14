package par_test

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"

	"github.com/piotrpersona/zero/par"
	"github.com/stretchr/testify/require"
)

func TestConcSlice(t *testing.T) {
	t.Parallel()

	t.Run("process concurrently", func(t *testing.T) {
		t.Parallel()

		slice := []int{1, 2, 3, 4, 5, 6}

		c := int64(0)

		err := par.Slice(slice, func(ctx context.Context, i int, val int) error {
			atomic.AddInt64(&c, 1)
			return nil
		})

		require.NoError(t, err)
		require.Equal(t, int64(6), atomic.LoadInt64(&c))
	})
	t.Run("exit on err", func(t *testing.T) {
		t.Parallel()

		slice := []int{1, 2, 3, 4, 5, 6}

		err := par.Slice(slice, func(ctx context.Context, i int, val int) error {
			if i%4 == 0 {
				return errors.New("err")
			}
			return nil
		})

		require.Error(t, err)
	})
	t.Run("handle ctx", func(t *testing.T) {
		t.Parallel()

		slice := []int{1, 2, 3, 4, 5, 6}

		ctx, cancel := context.WithCancel(context.Background())

		err := par.Slice(slice, func(ctx context.Context, i int, val int) error {
			if i%4 == 0 {
				cancel()
				return nil
			}
			return nil
		}, par.WithCtx(ctx))

		require.NoError(t, err)
	})
}
