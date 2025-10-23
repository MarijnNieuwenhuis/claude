package server

import (
	"github.com/gorilla/mux"
	"gitlab.com/btcdirect-api/bootstrap-go-service/internal/app"
	"gitlab.com/btcdirect-api/bootstrap-go-service/internal/http/handler"
)

// Registers all routes for the application.
func registerRoutes(r *mux.Router, app *app.App) {
	r.HandleFunc("/health", handler.HealthHandler(app)).Methods("GET")
	r.HandleFunc("/ready", handler.ReadinessHandler(app.DatabaseConnection())).Methods("GET")

	// TODO: Add your application-specific routes here
}
