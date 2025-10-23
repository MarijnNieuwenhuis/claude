package app

import (
	"context"
	"sync"
	"time"
)

// Contexts added to the graceful shutdown will be closed when a shutdown signal is received.
// In your application you can add listen to the context done (this also adds one to the wait groups)
// and call Done when the application is finished handling the shutdown.
//
// If you need more than one waitgroup, you can implement your own waitgroup in your service.
//
// Example for contexts:
//
//	func (a *App) Run() {
//		ctx, _ := a.Shutdown.Add()
//		go func() {
//			<-ctx.Done()
//			a.Shutdown.Done()
//		}()
//	}
type GracefulShutdown struct {
	cancels   []context.CancelFunc
	waitGroup sync.WaitGroup
}

func newGracefulShutdown() *GracefulShutdown {
	return &GracefulShutdown{
		cancels:   []context.CancelFunc{},
		waitGroup: sync.WaitGroup{},
	}
}

func (gs *GracefulShutdown) shutdown(timeout time.Duration) error {
	ctx, cancelCtx := context.WithTimeout(context.Background(), timeout)
	defer cancelCtx()

	go func() {
		for _, cancel := range gs.cancels {
			cancel()
		}

		gs.waitGroup.Wait()

		cancelCtx()
	}()

	<-ctx.Done()

	err := ctx.Err()
	if err == context.Canceled {
		return nil
	}

	return err
}

// Add a context to the graceful shutdown.
// This will also add one to the wait group.
func (gs *GracefulShutdown) Add() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	gs.cancels = append(gs.cancels, cancel)
	gs.waitGroup.Add(1)
	return ctx, cancel
}

// Done will remove one from the wait group.
func (gs *GracefulShutdown) Done() {
	gs.waitGroup.Done()
}
