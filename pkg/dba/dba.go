package dba

import (
	"database/sql"

	_ "github.com/lib/pq" // Import postgres db driver
	"github.com/sirupsen/logrus"
)

// DBA is an automatic database administrator.
type DBA struct {
	connStr string
	db      *sql.DB
	verbose bool
}

// New creates a new DBA and returns it
func New(connStr string) *DBA {
	return &DBA{connStr: connStr}
}

// Run runs the automatic DBA
func (dba *DBA) Run() error {
	var err error

	logrus.Info("Connecting to database")
	dba.db, err = sql.Open("postgres", dba.connStr)
	if err != nil {
		return err
	}
	defer dba.db.Close()

	if err = dba.db.Ping(); err != nil {
		return err
	}

	logrus.Info("Running Auto DBA")
	return dba.analyze()
}

func (dba *DBA) analyze() error {
	var cmd string
	if dba.verbose {
		cmd = "ANALYZE VERBOSE"
	} else {
		cmd = "ANALYZE"
	}

	_, err := dba.db.Exec(cmd)
	return err
}
