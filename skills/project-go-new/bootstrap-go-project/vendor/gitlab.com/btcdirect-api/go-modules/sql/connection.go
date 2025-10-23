package sql

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"cloud.google.com/go/cloudsqlconn"
	"cloud.google.com/go/cloudsqlconn/mysql/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type DBConnection interface {
	DB(autoRetry bool) *sqlx.DB
	IsAlive() bool
	Shutdown() error
}

// Connection is a wrapper around the sqlx.DB.
// Create the struct with all public fields and use the DB method to get the database connection.
type Connection struct {
	sync.Mutex
	Driver         string
	DSN            string
	Log            *zap.SugaredLogger
	ConnectTimeout time.Duration
	db             *sqlx.DB
}

type driver struct {
	Name    string
	Cleanup func() error
}

// DriverFromDSN determines the driver based on the DSN.
//
// Supported drivers:
// - mysql (default)
// - cloudsql-mysql (use the following DSN format: "myuser:mypass@cloudsql-mysql(project:region:instance)/mydb")
func DriverFromDSN(dsn string) (d driver, err error) {
	d.Name = "mysql"

	// CloudSQL MySQL
	if strings.Contains(dsn, "cloudsql-mysql") {
		d.Name = "cloudsql-mysql"
		d.Cleanup, err = mysql.RegisterDriver("cloudsql-mysql", cloudsqlconn.WithOptions(
			cloudsqlconn.WithIAMAuthN(),
			cloudsqlconn.WithDefaultDialOptions(
				cloudsqlconn.WithPrivateIP(),
			),
		))
	} else if strings.Contains(dsn, "sqlmock") {
		d.Name = "sqlmock"
		if strings.Contains(dsn, "cleanup=true") {
			d.Cleanup = func() error { return fmt.Errorf("test error") }
		}
	}

	return d, err
}

// Returns the database connection.
// If the connection is not yet established, it will try to establish the connection.
// If autoRetry is true, it will keep trying to establish the connection until it is successful.
func (c *Connection) DB(autoRetry bool) *sqlx.DB {
	if c.db == nil {
		c.setupDB(autoRetry)
	}

	return c.db
}

// Returns true if the database connection is alive.
// If the connection is not yet established, it will always return false.
func (c *Connection) IsAlive() bool {
	return c.db != nil && c.db.Ping() == nil
}

// Set up the database connection.
// If autoRetry is true, it will keep trying to establish the connection until it is successful.
//
// This method is thread-safe.
func (c *Connection) setupDB(autoRetry bool) {
	c.Lock()

	if c.db != nil {
		c.Unlock()
		return
	}

	db, err := sqlx.Open(c.Driver, c.DSN)

	if err == nil {
		err = db.Ping()
		if err == nil {
			c.Log.Info("Successfully connected to database")
			c.db = db
			c.Unlock()
			return
		}
	}

	c.Log.Errorf("Could not create database connection. %s", err.Error())

	if !autoRetry {
		c.Unlock()
		return
	}

	c.Log.Infof("Retrying to create database connection in %s...", c.ConnectTimeout.String())
	time.Sleep(c.ConnectTimeout)

	c.Unlock()
	c.setupDB(true)
}

// Close the database connection.
// If the connection is not yet established, it will do nothing.
//
// Will return an error if the database could not be closed.
//
// This method is thread-safe.
func (c *Connection) Shutdown() error {
	if c.db == nil {
		return nil
	}

	c.Lock()
	defer c.Unlock()

	c.Log.Info("Shutting down the database so we don't keep connections open")

	ctx, cancel := context.WithTimeout(context.Background(), c.ConnectTimeout)
	defer cancel()

	err := c.db.Close()
	if err != nil {
		c.Log.Infof("Could not close database %v", err.Error())
		return err
	}

	for {
		if err = c.db.Ping(); err.Error() == "sql: database is closed" {
			// Database is closed successfully.
			break
		}

		select {
		case <-ctx.Done():
			if err = ctx.Err(); err != nil {
				c.Log.Infof("Could not close database. %v", err.Error())
				return err
			}
			// Database is closed successfully.
			break
		default:
		}
	}

	c.Log.Info("Database shut down")

	return nil
}
