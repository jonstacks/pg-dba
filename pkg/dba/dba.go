package dba

import (
	"database/sql"

	"github.com/jonstacks/pg-dba/pkg/utils"
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
func New(connStr string, verbose bool) *DBA {
	return &DBA{
		connStr: connStr,
		verbose: verbose,
	}
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

	_, err := dba.exec(cmd)
	return err
}

func (dba *DBA) exec(query string, args ...interface{}) (sql.Result, error) {
	var err error
	var result sql.Result

	logrus.Infof("Running '%s'", query)
	runTime := utils.Time(func() { result, err = dba.db.Exec(query, args...) })
	context := logrus.WithFields(logrus.Fields{"duration": runTime.String()})
	if err != nil {
		context.Errorf("'%s' finished with error '%s'", query, err)
	} else {
		context.Infof("'%s' finished successfully", query)
	}
	return result, err
}
