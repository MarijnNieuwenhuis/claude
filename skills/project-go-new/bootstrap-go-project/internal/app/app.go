package app

import (
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/jmoiron/sqlx"
	"gitlab.com/btcdirect-api/go-modules/app"
	msg "gitlab.com/btcdirect-api/go-modules/messenger"
	"gitlab.com/btcdirect-api/go-modules/sql"
	"gitlab.com/btcdirect-api/go-modules/sql/migrate"
	"gitlab.com/btcdirect-api/bootstrap-go-service/internal/db"
	"go.uber.org/zap"
)

type App struct {
	config   Configuration
	database interface {
		Start() *sqlx.DB
		Connection() *sql.Connection
		Migrate(m migrate.Migrate) error
		Shutdown() error
	}
	messenger msg.Messenger
	handlers  []msg.MessageHandler
	core      *app.App
}

// Initialize the application.
// This will also load the configuration.
func Initialize(c Configuration) *App {
	// In development mode, we set the shutdown timeout to 0 to allow for instant shutdowns.
	// In production, we set it to 30 seconds to allow for graceful shutdowns.
	shutdownTimeout := 30 * time.Second
	if c.Environment == Dev {
		shutdownTimeout = 0
	}

	core := app.Initialize(
		app.WithLoggerForLevel(c.LogLevel),
		app.WithShutdownTimeout(shutdownTimeout),
	)

	database := db.New(c.DatabaseDSN, core.Log)
	database.Start()

	messenger := createMessenger(&core, c)

	// TODO: Add your message handlers here
	handlers := []msg.MessageHandler{}

	app := &App{
		config:    c,
		database:  database,
		messenger: messenger,
		handlers:  handlers,
		core:      &core,
	}

	app.initSentry()

	return app
}

// Run the application and its services.
func (a *App) Run() {
	for _, handler := range a.handlers {
		go a.messenger.Subscribe(handler)
	}

	a.core.Run()
}

// Migrate the database.
func (a *App) Migrate(m migrate.Migrate) error {
	return a.database.Migrate(m)
}

// Shutdown Shuts down all services of the application.
func (a *App) Shutdown() {
	if err := a.database.Shutdown(); err != nil {
		a.Logger().Errorf("error shutting down database: %v", err)
	}
	sentry.Flush(2 * time.Second)
}

// Config returns the application configuration.
func (a *App) Config() Configuration {
	return a.config
}

// Logger exposes the shared structured logger.
func (a *App) Logger() *zap.SugaredLogger {
	return a.core.Log
}

// DatabaseConnection exposes the database connection.
func (a *App) DatabaseConnection() *sql.Connection {
	return a.database.Connection()
}

func (a *App) initSentry() {
	if "" == a.config.SentryDSN {
		return
	}

	a.core.Log.Info("Starting to initialize Sentry - ", "DSN - ", a.config.SentryDSN)

	if err := sentry.Init(sentry.ClientOptions{
		Dsn:         a.config.SentryDSN,
		Environment: string(a.config.Environment),
	}); err != nil {
		a.core.Log.Panic("Failed to initialize Sentry", "error", err)
	}
}

func createMessenger(core *app.App, c Configuration) msg.Messenger {
	return msg.New(msg.Config{
		Log:            core.Log,
		Shutdown:       core.Shutdown,
		Environment:    string(c.Environment),
		RestartTimeout: 10 * time.Second,
		PubsubConfig: msg.PubsubConfig{
			Emulator:        c.Pubsub.Emulator,
			Project:         c.Pubsub.Project,
			DeadLetterTopic: "bootstrap-go-service.dead",
		},
	})
}
