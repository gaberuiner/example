package workerpool

import (
	"context"
)

// Simple represents a simple worker pool.
type Simple chan struct{}

// NewSimple creates a new Simple worker pool with the specified limit.
func NewSimple(limit int) Simple {
	return make(chan struct{}, limit)
}

// Exec executes the given work function in a worker goroutine within the worker pool.
// It returns an error if the context is canceled or if the worker pool is closed.
func (wp Simple) Exec(ctx context.Context, work func(ctx context.Context)) error {
	select {
	case <-ctx.Done():
		// Context is canceled, return error.
		return ctx.Err()
	case wp <- struct{}{}:
		// A worker is available, execute the work function in a goroutine.
		go func() {
			work(ctx)
			select {
			case <-ctx.Done():
				// Context is canceled, release the worker.
				<-wp
			case <-wp:
				// Worker is released.
			}
		}()
	}

	return nil
}
