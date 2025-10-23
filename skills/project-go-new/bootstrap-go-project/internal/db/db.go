package db

import (
	"embed"
	"time"

	"github.com/jmoiron/sqlx"
	"gitlab.com/btcdirect-api/go-modules/sql"
	"gitlab.com/btcdirect-api/go-modules/sql/migrate"
	"go.uber.org/zap"
)

type database struct {
	log           *zap.SugaredLogger
	conn          *sql.Connection
	driverCleanup func() error
}

//go:embed migrations/*
var migrations embed.FS

// New creates a new database instance.
// The database Connection is not yet established, use the Start method to do that.
// The DSN is used to connect to the database, an error is returned if the DSN is invalid.
//
// Cloud SQL is supported by using the following DSN format: "myuser:mypass@cloudsql-mysql(project:region:instance)/mydb"
func New(dsn string, log *zap.SugaredLogger) *database {
	l := log.With("component", "database")
	d, _ := sql.DriverFromDSN(dsn)

	conn := &sql.Connection{
		Driver:         d.Name,
		DSN:            dsn,
		Log:            l,
		ConnectTimeout: 10 * time.Second,
	}

	return &database{
		log:           l,
		conn:          conn,
		driverCleanup: d.Cleanup,
	}
}

// Start opens the Connection to the database.
// This will block until the Connection is established.
// This should be called once during application startup.
func (db *database) Start() *sqlx.DB {
	db.log.Info("Connecting to the database")

	return db.conn.DB(true)
}

// Migrate the database.
func (db *database) Migrate(m migrate.Migrate) error {
	return m.Migrate(migrations, db.conn, db.log)
}

// Shutdown closes the database Connection and cleans up the driver if needed.
func (db *database) Shutdown() error {
	if err := db.conn.Shutdown(); err != nil {
		return err
	}

	if db.driverCleanup != nil {
		return db.driverCleanup()
	}

	return nil
}

// Connection returns the database connection.
func (db *database) Connection() *sql.Connection {
	return db.conn
}
