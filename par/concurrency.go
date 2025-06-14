package par

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type parallelConfig struct {
	limit int
	ctx   context.Context
}

type ParallelOption func(*parallelConfig)

func WithLimit(limit int) ParallelOption {
	return func(pc *parallelConfig) {
		pc.limit = limit
	}
}

func WithCtx(ctx context.Context) ParallelOption {
	return func(pc *parallelConfig) {
		pc.ctx = ctx
	}
}

func applyOpts(opts ...ParallelOption) *parallelConfig {
	cfg := parallelConfig{
		limit: 16,
		ctx:   context.Background(),
	}
	for _, opt := range opts {
		opt(&cfg)
	}
	return &cfg
}

func IterSlice[T any](arr []T, callback func(context.Context, int, T) error, opts ...ParallelOption) error {
	cfg := applyOpts(opts...)

	pool, gctx := errgroup.WithContext(cfg.ctx)
	pool.SetLimit(cfg.limit)

	for index, elem := range arr {
		pool.Go(func() error {
			return callback(gctx, index, elem)
		})
	}
	return pool.Wait()
}

func IterMap[K comparable, V any](hashMap map[K]V, callback func(context.Context, K, V) error, opts ...ParallelOption) error {
	cfg := applyOpts(opts...)

	pool, gctx := errgroup.WithContext(cfg.ctx)
	pool.SetLimit(cfg.limit)

	for key, value := range hashMap {
		pool.Go(func() error {
			return callback(gctx, key, value)
		})
	}
	return pool.Wait()
}
