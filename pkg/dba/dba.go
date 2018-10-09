package dba

import (
	"database/sql"

	_ "github.com/lib/pq" // Import postgres db driver
)

// DBA is an automatic database administrator.
type DBA struct {
	db *sql.DB
}

// New creates a new DBA and returns it
func New(connStr string) (*DBA, error) {
	var err error
	dba := &DBA{}

	dba.db, err = sql.Open("postgres", connStr)
	return dba, err
}

// Run runs the automatic DBA
func (dba *DBA) Run() error {
	return dba.db.Ping()
}
