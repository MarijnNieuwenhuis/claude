package migrate

import (
	"embed"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jmoiron/sqlx"
	"gitlab.com/btcdirect-api/go-modules/sql"
	"go.uber.org/zap"
)

const maxDatabaseAttempts = 10

type Migrate struct {
	Cmd, Param string
}

type migration struct {
	Cmd, Param string
	Migrate    *migrate.Migrate
	Log        *zap.SugaredLogger
}

// Migrate is a function that runs the migrations for the given connection.
// The migrations are loaded from the given filesystem.
// The filesystem should contain a directory called 'migrations' with the migration files.
//
// Supported commands are:
//
//   - up: Perform all up migrations
//   - down: Perform all down migrations
//   - version: Print the current migration version
//   - force: Force the migration version to the given version
//   - target: Migrate to the given version
//   - steps: Perform the given number of migration steps
//
// The Param field is used as the version for the force, target and steps commands.
func (m Migrate) Migrate(fs embed.FS, conn *sql.Connection, log *zap.SugaredLogger) error {
	log.Info("Running database migrations")
	defer log.Info("Finished running database migrations")

	mi, err := createMigrateInstance(fs, conn, log)
	if err != nil {
		return err
	}

	migration := &migration{
		Cmd:     m.Cmd,
		Param:   m.Param,
		Migrate: mi,
		Log:     log,
	}

	return migration.Run()
}

// Wrapper for running the golang-migrate/migrate/v4 package.
func (m *migration) Run() (err error) {
	switch m.Cmd {
	case "":
		fallthrough
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	case "version":
		err = m.Version()
	case "force":
		err = m.Force()
	case "target":
		err = m.Target()
	case "steps":
		err = m.Steps()
	}

	if err == nil {
		return
	}

	if err.Error() == "no change" {
		m.Log.Info("No changes were made. Target version is the same as current version")
		return nil
	}

	m.Log.Errorf("Migration failed with error: %v", err.Error())
	return err
}

func (m *migration) Up() error {
	m.Log.Info("Performing all up migrations")
	return m.Migrate.Up()
}

func (m *migration) Down() error {
	m.Log.Info("Performing all down migrations")
	return m.Migrate.Down()
}

func (m *migration) Version() error {
	v, d, err := m.Migrate.Version()
	if err != nil {
		return fmt.Errorf("Error retrieving current migration version: %s", err.Error())
	}
	state := map[bool]string{true: "DIRTY", false: "CLEAN"}[d]
	m.Log.Infof("Current migration version is '%d' (%s)", v, state)
	return nil
}

func (m *migration) Force() error {
	i, err := m.intParam()
	if err != nil {
		return err
	}
	m.Log.Infof("Forcing migration version to '%d'", i)
	return m.Migrate.Force(i)
}

func (m *migration) Target() error {
	i, err := m.intParam()
	if err != nil {
		return err
	}
	m.Log.Infof("Migration version to '%d'", i)
	return m.Migrate.Migrate(uint(i))
}

func (m *migration) Steps() error {
	i, err := m.intParam()
	if err != nil {
		return err
	}
	m.Log.Infof("Performing '%d' migration steps", i)
	return m.Migrate.Steps(i)
}

func (m *migration) intParam() (int, error) {
	i, err := strconv.Atoi(m.Param)
	if err != nil {
		return 0, fmt.Errorf("Argument '%s' is not a valid migration integer", m.Param)
	}
	return i, nil
}

// Retrieves the database from the given connection.
// It will wait for the database to become available for a maximum of 10 seconds.
// If the database is not available after 10 seconds, an error is returned.
func database(conn *sql.Connection, log *zap.SugaredLogger) (*sqlx.DB, error) {
	for attempt := 1; attempt <= maxDatabaseAttempts; attempt++ {
		conn.DB(false)
		if conn.IsAlive() {
			break
		}

		log.Info("Waiting for database to become available")
		time.Sleep(1 * time.Second)
	}

	if !conn.IsAlive() {
		err := fmt.Errorf("Giving up after %d attempts to connect to database", maxDatabaseAttempts)
		log.Info(err.Error())
		return nil, err
	}

	return conn.DB(false), nil
}

// Creates a new migrate instance with the given filesystem, connection and logger.
//
// The filesystem should contain a directory called 'migrations' with the migration files.
func createMigrateInstance(fs embed.FS, conn *sql.Connection, log *zap.SugaredLogger) (m *migrate.Migrate, err error) {
	db, err := database(conn, log)
	if err != nil {
		return
	}

	driver, err := mysql.WithInstance(db.DB, &mysql.Config{})
	if err != nil {
		return
	}

	d, err := iofs.New(fs, "migrations")
	if err != nil {
		return
	}

	m, err = migrate.NewWithInstance(
		"iofs", d,
		conn.Driver, driver)

	return
}
