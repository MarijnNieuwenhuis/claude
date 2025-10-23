package app

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/coreos/go-systemd/daemon"
	"gitlab.com/btcdirect-api/go-modules/logger"
	"go.uber.org/zap"
)

// App struct, you should embed this in your own application struct
// to add custom services.
//
// Example:
//
//	type App struct {
//		*app.App
//		MyService *service.MyService
//	}
type App struct {
	Log             *zap.SugaredLogger
	Shutdown        *GracefulShutdown
	shutdownTimeout time.Duration
}

type opt func(*App)

// Initialize creates an application and applies the given options.
func Initialize(opts ...opt) App {
	a := App{
		Shutdown: newGracefulShutdown(),
	}

	for _, o := range opts {
		o(&a)
	}

	return a
}

// WithLogger sets the logger for the application.
func WithLogger(log *zap.SugaredLogger) opt {
	return func(a *App) {
		a.Log = log
	}
}

// WithLoggerForLevel creates a logger for the given log level and sets it for the application.
func WithLoggerForLevel(logLevel string) opt {
	return func(a *App) {
		a.Log = logger.NewLogger(logLevel)
	}
}

// WithShutdownTimeout sets a timeout to wait before shutting down the application.
// This can be useful for a graceful shutdown in Kubernetes as it cannot use a preStop hook due to
// the container being distroless.
func WithShutdownTimeout(timeout time.Duration) opt {
	return func(a *App) {
		a.shutdownTimeout = timeout
	}
}

// Run the application, this will block until a shutdown signal is received.
// This will also notify systemd that the application is ready.
//
// When a shutdown signal is received, all stop channels will be closed aswell.
func (a *App) Run() {
	if runtime.GOOS == "linux" {
		// Notify systemd that the application is ready.
		daemon.SdNotify(false, "READY=1")
	}

	a.waitForShutdown()

	if a.shutdownTimeout > 0 {
		if a.Log != nil {
			a.Log.Infof("Waiting %s before shutting down application...", a.shutdownTimeout)
		}
		time.Sleep(a.shutdownTimeout)
	}

	if err := a.Shutdown.shutdown(30 * time.Second); err != nil {
		a.Log.Error(err)
	}
}

func (a *App) waitForShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	for {
		// This will block the process until a shutdown signal is received.
		switch <-c {
		case syscall.SIGINT, syscall.SIGTERM:
			if a.Log != nil {
				a.Log.Info("Shutdown request received.")
			}
			return
		}
	}
}
