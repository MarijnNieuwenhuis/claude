package http

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// Server is a wrapper around the http.Server.
type server struct {
	Router *mux.Router
	server *http.Server
	log    *zap.SugaredLogger
}

// CreateServer creates a new HTTP server with the given port and logger.
// The logger will be used to log the HTTP requests.
//
// Add your own routes to the router and start the server with the Start method.
func CreateServer(port string, log *zap.SugaredLogger) server {
	r := mux.NewRouter()
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: createLoggingRouter(r, log),
	}
	s := server{
		Router: r,
		server: srv,
		log:    log,
	}

	return s
}

// Start the HTTP server.
func (s server) Start() {
	s.log.Infof("Starting HTTP server on %s", s.server.Addr)

	go s.run()
}

// Run the HTTP server, this will block until the server is shutdown.
func (s server) run() {
	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		s.log.Fatalf("Failed to start HTTP server: %s", err)
	}
}

// Gracefully shutdown the HTTP server.
// If the server is not shutdown within 5 seconds, the server will be forcefully shutdown.
func (s server) Shutdown() {
	s.log.Info("Shutting down HTTP server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		s.log.Fatalf("Failed to shutdown HTTP server: %s", err)
	}

	s.log.Info("HTTP server shutdown")
}
