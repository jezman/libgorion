package libgorion

import (
	"database/sql"
	"time"
)

// Datastore methods
type Datastore interface {
	AddWorker(string) error
	DeleteWorker(string) error
	DisableWorkerCard(string) error
	EnableWorkerCard(string) error
	Company() ([]*Company, error)
	Doors() ([]*Door, error)
	Workers(string) ([]*Worker, error)
	Events(string, string, string, uint, bool) ([]*Event, error)
	EventsValues() ([]*Event, error)
	EventsTail(time.Duration, string) error
	WorkedTime(string, string, string, string) ([]*Event, error)
}

// Database structure used as receiver in methods
type Database struct {
	*sql.DB
}

// OpenDB opening connecting to server
// return pointer to struct DB and error
func OpenDB(dsn string) (*Database, error) {
	db, err := sql.Open("mssql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &Database{DB: db}, nil
}
