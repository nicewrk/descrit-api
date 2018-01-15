package store

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"regexp"

	// Necessary for establishing the DB connection.
	_ "github.com/lib/pq"
)

const (
	connStringFormat = "postgres://%s:%s@%s:%s/%s"
	driver           = "postgres"
)

// API represents the interface for interacting with the database,
// which will be implemented by Client.
type API interface {
	Prepare(string) (*sql.Stmt, *sql.Tx, error)
	Execute(*sql.Stmt, *sql.Tx, ...interface{}) error
	Rollback(*sql.Tx, error) error
}

// Client implements API.
type Client struct {
	db *sql.DB
}

// NewClient configures and returns a store (DB) client.
func NewClient() (*Client, error) {
	db, err := sql.Open(driver, connURI())
	if err != nil {
		return nil, err
	}
	// Check ability to make a DB connection.
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Client{db: db}, nil
}

// Prepare BEGINs a DB transaction and PREPAREs a DB statement.
func (c *Client) Prepare(SQL string) (*sql.Stmt, *sql.Tx, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return nil, nil, err
	}
	stmt, err := tx.Prepare(SQL)
	return stmt, tx, err
}

// Execute EXECUTEs a DB statement.
func (c *Client) Execute(stmt *sql.Stmt, tx *sql.Tx, args ...interface{}) error {
	defer func() {
		err := stmt.Close()
		if err != nil {
			log.Printf("error: closing stmt | %s", err)
		}
	}()

	_, err := stmt.Exec(args...)
	if err != nil {
		return c.Rollback(tx, err)
	}
	return tx.Commit()
}

// Rollback ROLLSBACK a DB transaction.
func (c *Client) Rollback(tx *sql.Tx, err error) error {
	e := tx.Rollback()
	if e != nil {
		log.Printf("error: rolling back transaction: %s", e)
	}
	return err
}

func connURI() string {
	connURI := fmt.Sprintf(connStringFormat, os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	isProductionOrStaging, err := regexp.MatchString("(production|staging)", os.Getenv("ENVIRONMENT"))
	if err != nil {
		log.Fatalf("error: identifying environment: %s.", err)
	}
	if !isProductionOrStaging {
		connURI += "?sslmode=disable"
	}

	return connURI
}
