package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jonstacks/pg-dba/pkg/config"
	"github.com/jonstacks/pg-dba/pkg/dba"
	"github.com/sirupsen/logrus"
)

func fatal(err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
}

func init() {
	logrus.SetLevel(config.LogLevel())
	logrus.SetFormatter(config.LogFormat())
}

func main() {

	analyzeTimeoutSecs, err := config.AnalyzeTimeoutSeconds()
	fatal(err)

	vacuumTimeoutSecs, err := config.VacuumTimeoutSeconds()
	fatal(err)

	fullVacuumTimeoutSecs, err := config.FullVacuumTimeoutSeconds()
	fatal(err)

	opts := dba.NewDefaultOptions()
	opts.AnalyzeTimeout = time.Duration(analyzeTimeoutSecs) * time.Second
	opts.VacuumTimout = time.Duration(vacuumTimeoutSecs) * time.Second
	opts.FullVacuumTimeout = time.Duration(fullVacuumTimeoutSecs)
	opts.Verbose = config.Verbose()
	admin := dba.New(config.DBConnectionString(), opts)
	fatal(admin.Run())
}
