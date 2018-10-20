package dba

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jonstacks/pg-dba/pkg/utils"
	_ "github.com/lib/pq" // Import postgres db driver
	"github.com/sirupsen/logrus"
)

// Options controll how the DBA determines what to do
type Options struct {
	PreAnalyze     bool
	AnalyzeTimeout time.Duration
	PostAnalyze    bool

	PreVacuum    bool
	VacuumTimout time.Duration

	FullVacuumBloatPercent   int           // Only consider tables with bloat larger than this percent to need a full vacuum
	FullVacuumMaxTableSizeMb int           // Do not vacuum tables larger than this size due to the amount of time it would take
	FullVacuumTimeout        time.Duration // Cancel full vacuum's that take longer than this many seconds

	Timeout time.Duration // Default timeout for everything else
	Verbose bool
}

// NewDefaultOptions returns new Options initilized with our reasonable defaults
func NewDefaultOptions() Options {
	return Options{
		PreAnalyze:     true,
		AnalyzeTimeout: 10 * time.Minute,
		PostAnalyze:    true,

		PreVacuum:    false,
		VacuumTimout: 10 * time.Minute,

		FullVacuumBloatPercent:   10,
		FullVacuumMaxTableSizeMb: 10 * 1000, // Only Vacuum tables smaller than 10 GB by default
		FullVacuumTimeout:        10 * time.Minute,

		Timeout: 30 * time.Second,
		Verbose: false,
	}
}

// DBA is an automatic database administrator.
type DBA struct {
	connStr string
	db      *sql.DB
	options Options
}

// New creates a new DBA and returns it
func New(connStr string, opts Options) *DBA {
	return &DBA{
		connStr: connStr,
		options: opts,
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

	if dba.options.PreVacuum {
		logrus.Info("Running a vacuum to clean up space")
		if err = dba.vacuum(dba.options.PreAnalyze); err != nil {
			return err
		}
	} else if dba.options.PreAnalyze {
		logrus.Info("Running pre-analyze")
		if err = dba.analyze(); err != nil {
			return err
		}
	}

	// Now let's see which tables need a full vacuum
	results, err := dba.bloatedTables()
	if err != nil {
		return err
	}

	for _, r := range results {
		logContext := logrus.WithFields(logrus.Fields{
			"table_mb":      r.TableMb,
			"table_name":    r.TableName,
			"schema_name":   r.SchemaName,
			"percent_bloat": r.PercentBloat,
		})
		if r.TableMb <= float64(dba.options.FullVacuumMaxTableSizeMb) {
			logContext.Info("Table can be vacuumed.")
			err = dba.fullVacuum(fmt.Sprintf("%s.%s", r.SchemaName, r.TableName))
			if err != nil {
				return err
			}
		} else {
			logContext.Warn("Table can NOT be vacuumed due to table size.")
		}
	}

	if dba.options.PostAnalyze {
		logrus.Info("Running post-analyze to update statistics for query planner")
		err = dba.analyze()
	}

	return err
}

func (dba *DBA) analyze() error {
	var cmd string
	if dba.options.Verbose {
		cmd = "ANALYZE VERBOSE"
	} else {
		cmd = "ANALYZE"
	}

	ctx, cancel := context.WithTimeout(context.Background(), dba.options.AnalyzeTimeout)
	defer cancel()

	_, err := dba.execContext(ctx, cmd)
	return err
}

func (dba *DBA) fullVacuum(tableName string) error {
	cmd := fmt.Sprintf("VACUUM (FULL, ANALYZE) %s", tableName)
	ctx, cancel := context.WithTimeout(context.Background(), dba.options.VacuumTimout)
	defer cancel()
	_, err := dba.execContext(ctx, cmd)
	return err
}

func (dba *DBA) vacuum(analyze bool) error {
	var cmd string
	if dba.options.Verbose {
		cmd = "VACUUM (VERBOSE, ANALYZE)"
	} else {
		cmd = "VACUUM ANALYZE"
	}

	ctx, cancel := context.WithTimeout(context.Background(), dba.options.VacuumTimout)
	defer cancel()
	_, err := dba.execContext(ctx, cmd)
	return err
}

func (dba *DBA) bloatedTables() ([]TableBloatResult, error) {
	results := make([]TableBloatResult, 0)
	query := TableBloatQuery(dba.options.FullVacuumBloatPercent)
	ctx, cancel := context.WithTimeout(context.Background(), dba.options.Timeout)
	defer cancel()
	rows, err := dba.queryContext(ctx, query)
	if err != nil {
		return results, err
	}
	for rows.Next() {
		var r TableBloatResult
		if err := rows.Scan(&r.DatabaseName, &r.SchemaName, &r.TableName, &r.CanEstimate, &r.EstimatedRows, &r.PercentBloat, &r.MbBloat, &r.TableMb); err != nil {
			return results, err
		}
		results = append(results, r)
	}
	return results, nil
}

func (dba *DBA) queryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	var err error
	var rows *sql.Rows

	logrus.Debugf("Running '%s'", query)
	runTime := utils.Time(func() { rows, err = dba.db.QueryContext(ctx, query, args...) })
	logContext := logrus.WithFields(logrus.Fields{"duration": runTime.String()})
	if err != nil {
		logContext.Errorf("'%s' finished with error '%s'", query, err)
	} else {
		logContext.Infof("'%s' finished successfully", query)
	}
	return rows, err
}

func (dba *DBA) execContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	var err error
	var result sql.Result

	logrus.Debugf("Running '%s'", query)
	runTime := utils.Time(func() { result, err = dba.db.ExecContext(ctx, query, args...) })
	logContext := logrus.WithFields(logrus.Fields{"duration": runTime.String()})
	if err != nil {
		logContext.Errorf("'%s' finished with error '%s'", query, err)
	} else {
		logContext.Infof("'%s' finished successfully", query)
	}
	return result, err
}
