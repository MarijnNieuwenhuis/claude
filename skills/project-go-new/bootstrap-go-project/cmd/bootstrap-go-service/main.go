package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"gitlab.com/btcdirect-api/go-modules/sql/migrate"
	"gitlab.com/btcdirect-api/bootstrap-go-service/internal/app"
	"gitlab.com/btcdirect-api/bootstrap-go-service/internal/http/server"
)

func main() {
	c := app.Configuration{}

	var env string
	flag.StringVar(&env, "env", getenv("APP_ENV", "dev"), "Environment")

	var err error
	c.Environment, err = getEnvironment(env)
	if err != nil {
		panic(err)
	}

	flag.StringVar(&c.LogLevel, "loglevel", getenv("LOG_LEVEL", "info"), "Log output level")
	flag.StringVar(&c.HTTPPort, "port", getenv("HTTP_PORT", "8080"), "HTTP port")
	flag.StringVar(&c.DatabaseDSN, "database", os.Getenv("DATABASE_URL"), "Database dsn")
	flag.StringVar(&c.SentryDSN, "sentry-dsn", os.Getenv("SENTRY_DSN"), "Sentry DSN")

	flag.StringVar(&c.Pubsub.Emulator, "pubsub-emulator", os.Getenv("PUBSUB_EMULATOR"), "Pubsub emulator host")
	flag.StringVar(&c.Pubsub.Project, "pubsub-project", os.Getenv("PUBSUB_PROJECT"), "Pubsub project id")

	var migrate bool
	flag.BoolVar(&migrate, "migrate", false, "Run database migrations")

	flag.Parse()

	if migrate {
		// Allow multi statement for migrations.
		suffix := "?"
		if strings.Contains(c.DatabaseDSN, suffix) {
			suffix = "&"
		}
		c.DatabaseDSN += suffix + "multiStatements=true"
	}

	application := app.Initialize(c)

	if migrate {
		migr(application)
	} else {
		run(application)
	}
}

// Run the application in migrate mode.
func migr(application *app.App) {
	m := migrate.ParseMigrationFlags("migrate")
	if err := application.Migrate(m); err != nil {
		application.Logger().Errorf("Error migrating: %v", err)
		os.Exit(1)
	}

	os.Exit(0)
}

// Run the application daemon.
func run(application *app.App) {
	application.Logger().Info("Starting application")

	server := server.Start(application)
	application.Run()

	application.Logger().Info("Shutting down application")

	application.Shutdown()
	server.Shutdown()

	os.Exit(0)
}

func getenv(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func getEnvironment(input string) (app.Environment, error) {
	switch input {
	case "dev":
		return app.Dev, nil
	case "stage":
		return app.Stage, nil
	case "acc":
		return app.Acc, nil
	case "sandbox":
		return app.Sandbox, nil
	case "prod":
		return app.Prod, nil
	default:
		return "", fmt.Errorf("invalid environment: %s", input)
	}
}
