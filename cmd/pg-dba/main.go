package main

import (
	"fmt"
	"os"

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
	opts := dba.NewDefaultOptions()
	opts.Verbose = config.Verbose()
	admin := dba.New(config.DBConnectionString(), opts)
	fatal(admin.Run())
}
